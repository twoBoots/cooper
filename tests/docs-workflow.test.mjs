import { test } from 'node:test';
import assert from 'node:assert/strict';
import fs from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const repoRoot = path.resolve(__dirname, '..');

test('GitHub Actions workflow .github/workflows/pages.yml exists and has proper configuration', () => {
  const workflowPath = path.join(repoRoot, '.github', 'workflows', 'pages.yml');
  assert.ok(fs.existsSync(workflowPath), '.github/workflows/pages.yml must exist');

  const content = fs.readFileSync(workflowPath, 'utf8');

  // Triggers
  assert.ok(content.includes('branches: [main]') || content.includes('- main'), 'workflow must trigger on push to main');
  assert.ok(content.includes('workflow_dispatch'), 'workflow must allow manual dispatch');

  // Permissions
  assert.ok(content.includes('pages: write'), 'must have pages: write permission');
  assert.ok(content.includes('id-token: write'), 'must have id-token: write permission');

  // Build & Deploy steps
  assert.ok(content.includes('actions/configure-pages'), 'must use actions/configure-pages');
  assert.ok(content.includes('actions/upload-pages-artifact'), 'must use actions/upload-pages-artifact');
  assert.ok(content.includes('actions/deploy-pages'), 'must use actions/deploy-pages');
  assert.ok(content.includes('docs/.vitepress/dist'), 'artifact path must be docs/.vitepress/dist');
  assert.ok(content.includes('npm ci'), 'must use npm ci for reproducible build');
  assert.ok(content.includes('npm run docs:build'), 'must execute npm run docs:build');
});

test('Workflow policy, starter template, and documentation spec enforce Living Files Stand Alone', () => {
  const workflowPath = path.join(repoRoot, '.cooper', 'definition', 'workflow.md');
  const templatePath = path.join(repoRoot, 'templates', 'workflow.md');
  const docSpecPath = path.join(repoRoot, '.cooper', 'specs', 'documentation', 'spec.md');

  for (const [label, filePath] of [
    ['.cooper/definition/workflow.md', workflowPath],
    ['templates/workflow.md', templatePath],
  ]) {
    const text = fs.readFileSync(filePath, 'utf8');
    assert.ok(text.includes('**Living Files Stand Alone:**'), `${label} must define Living Files Stand Alone principle`);
    assert.ok(text.includes('.cooper/definition/') && text.includes('.cooper/specs/'), `${label} must scope rule to definition/ and specs/`);
    assert.ok(text.includes('.cooper/active/') && text.includes('.cooper/archive/'), `${label} must exempt active/ and archive/`);
    assert.ok(text.includes('Rationale Links'), `${label} must document the rationale-link exception`);
    assert.ok(text.includes('Open Known Gaps'), `${label} must document the open known-gap exception`);
    assert.ok(text.includes('stripping any planning provenance'), `${label} must require stripping planning provenance during Track Finalization`);
  }

  const specText = fs.readFileSync(docSpecPath, 'utf8');
  assert.ok(specText.includes('### Requirement: Living Files Stand Alone'), 'documentation/spec.md must include Living Files Stand Alone requirement');
  assert.ok(specText.includes('#### Scenario: Living files reject ticket, decision-marker, and bare section-number references'), 'must include core rejection scenario');
  assert.ok(specText.includes('#### Scenario: Rationale links outside requirement text cite frozen archived tracks or approved RFCs by file path and heading name'), 'must include rationale link scenario');
  assert.ok(specText.includes('#### Scenario: Open known-gap entries outside requirement text cite tracking tickets until resolved'), 'must include known-gap scenario');
  assert.ok(specText.includes('#### Scenario: Spec delta finalization strips planning provenance before merging into living specs'), 'must include finalization cleanup scenario');
  assert.ok(specText.includes('#### Scenario: Brownfield audit flags forbidden references in existing living files'), 'must include brownfield audit scenario');
});

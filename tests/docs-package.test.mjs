import { test } from 'node:test';
import assert from 'node:assert/strict';
import fs from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const repoRoot = path.resolve(__dirname, '..');

test('package.json exists and defines required documentation scripts', () => {
  const pkgPath = path.join(repoRoot, 'package.json');
  assert.ok(fs.existsSync(pkgPath), 'package.json must exist in repo root');

  const pkg = JSON.parse(fs.readFileSync(pkgPath, 'utf8'));
  assert.equal(pkg.name, 'cooper-docs', 'package name should be cooper-docs');
  assert.ok(pkg.scripts, 'package.json must have scripts');
  assert.ok(pkg.scripts['docs:dev'], 'scripts must define docs:dev');
  assert.ok(pkg.scripts['docs:build'], 'scripts must define docs:build');
  assert.ok(pkg.scripts['docs:preview'], 'scripts must define docs:preview');
  assert.ok(pkg.scripts.test, 'scripts must define test');

  const devDeps = pkg.devDependencies || {};
  assert.ok(devDeps.vitepress, 'vitepress must be listed in devDependencies');
});

test('package-lock.json exists ensuring reproducible builds', () => {
  const lockPath = path.join(repoRoot, 'package-lock.json');
  assert.ok(fs.existsSync(lockPath), 'package-lock.json must exist');
});

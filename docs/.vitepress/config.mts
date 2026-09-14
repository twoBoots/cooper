import { defineConfig } from 'vitepress';

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: 'Cooper',
  description: 'Spec-Driven Development (SDD) & Troop Worktree Isolation Framework',
  base: process.env.VITEPRESS_BASE || '/cooper/',
  themeConfig: {
    siteTitle: 'Cooper',
    nav: [
      { text: 'Guide', link: '/guide/getting-started' },
      { text: 'Workflow', link: '/guide/workflow' },
      { text: 'Installation', link: '/INSTALL' },
      { text: 'Comparisons', link: '/openspec-vs-conductor-comparison' }
    ],
    sidebar: [
      {
        text: 'Introduction',
        items: [
          { text: 'Getting Started', link: '/guide/getting-started' },
          { text: 'Workflow & Lifecycle', link: '/guide/workflow' },
          { text: 'Installation Guide', link: '/INSTALL' }
        ]
      },
      {
        text: 'Architecture & RFCs',
        items: [
          { text: 'Draft PRs in RFCs', link: '/rfc-draft-prs' },
          { text: 'Comparison vs Conductor', link: '/openspec-vs-conductor-comparison' }
        ]
      }
    ],
    socialLinks: [
      { icon: 'github', link: 'https://github.com/twoBoots/cooper' }
    ],
    footer: {
      message: 'Released under the MIT License.',
      copyright: 'Copyright © 2026 twoBoots'
    },
    search: {
      provider: 'local'
    }
  }
});

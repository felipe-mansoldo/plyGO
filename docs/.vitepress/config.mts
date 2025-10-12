import { defineConfig } from 'vitepress'

export default defineConfig({
  title: "plyGO",
  description: "Data Manipulation in Go, Simplified",
  base: '/plyGO/',

  themeConfig: {
    logo: '/img/logo.svg',

    nav: [
      { text: 'Home', link: '/' },
      { text: 'Guide', link: '/guide/getting-started' },
      { text: 'Basics', link: '/basics/data-loading' },
      { text: 'Advanced', link: '/advanced/composition' },
      { text: 'Extras', link: '/extras/csv-loading' }
    ],

    sidebar: [
      {
        text: 'Getting Started',
        items: [
          { text: 'Introduction', link: '/guide/intro' },
          { text: 'Getting Started', link: '/guide/getting-started' }
        ]
      },
      {
        text: 'Basics',
        items: [
          { text: 'Data Loading', link: '/basics/data-loading' },
          { text: 'Filtering', link: '/basics/filtering' },
          { text: 'Selecting', link: '/basics/selecting' },
          { text: 'Positions', link: '/basics/positions' },
          { text: 'Sorting', link: '/basics/sorting' },
          { text: 'Grouping', link: '/basics/grouping' },
          { text: 'Transformation', link: '/basics/transformation' },
          { text: 'Utilities', link: '/basics/utilities' },
          { text: 'Show', link: '/basics/show' },
          { text: 'Error Handling', link: '/basics/error-handling' }
        ]
      },
      {
        text: 'Advanced',
        items: [
          { text: 'Composition', link: '/advanced/composition' },
          { text: 'Custom Helpers', link: '/advanced/custom-helpers' },
          { text: 'Performance', link: '/advanced/performance' },
          { text: 'Concurrency', link: '/advanced/concurrency' },
          { text: 'Large Data', link: '/advanced/large-data' }
        ]
      },
      {
        text: 'Extras',
        items: [
          { text: 'CSV Loading', link: '/extras/csv-loading' },
          { text: 'Real World Examples', link: '/extras/real-world-examples' },
          { text: 'FAQ', link: '/extras/faq' }
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/mansoldof/plyGO' }
    ],

    search: {
      provider: 'local'
    },

    editLink: {
      pattern: 'https://github.com/mansoldof/plyGO/edit/main/docs/:path',
      text: 'Edit this page on GitHub'
    },

    footer: {
      message: 'Released under the MIT License.',
      copyright: 'Copyright © 2024-present'
    }
  },

  markdown: {
    theme: {
      light: 'github-light',
      dark: 'github-dark'
    },
    lineNumbers: true
  }
})

module.exports = {
  tutorialSidebar: [
    'intro',
    {
      type: 'category',
      label: '1️⃣ Getting Started',
      collapsed: false,
      items: ['tutorial-basics/getting-started'],
    },
    {
      type: 'category',
      label: '2️⃣ Basics',
      items: [
        'tutorial-basics/data-loading',
        'tutorial-basics/filtering',
        'tutorial-basics/selecting',
        'tutorial-basics/sorting',
        'tutorial-basics/grouping',
        'tutorial-basics/transformation',
        'tutorial-basics/positions',
        'tutorial-basics/utilities',
        'tutorial-basics/show',
        'tutorial-basics/error-handling',
      ],
    },
    {
      type: 'category',
      label: '3️⃣ Advanced',
      items: [
        'tutorial-advanced/composition',
        'tutorial-advanced/custom-helpers',
        'tutorial-advanced/performance',
        'tutorial-advanced/concurrency',
        'tutorial-advanced/large-data',
      ],
    },
    {
      type: 'category',
      label: '4️⃣ Extras',
      items: [
        'extras/csv-loading',
        'extras/real-world-examples',
        'extras/faq',
      ],
    },
  ],
};

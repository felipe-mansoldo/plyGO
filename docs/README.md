# plyGO Documentation Site

This directory contains the Docusaurus documentation site for plyGO.

## 🚀 Quick Start

### Install Dependencies

```bash
npm install
```

### Local Development

```bash
npm start
```

This command starts a local development server and opens up a browser window. Most changes are reflected live without having to restart the server.

### Build

```bash
npm run build
```

This command generates static content into the `build` directory and can be served using any static contents hosting service.

### Serve Built Site

```bash
npm run serve
```

Test the production build locally.

## 📁 Project Structure

```
docs/
├── docs/                    # Documentation pages
│   ├── intro.md            # Introduction
│   ├── tutorial-basics/    # Basic tutorials
│   ├── tutorial-advanced/  # Advanced topics
│   ├── extras/             # Extra content
│   └── api/                # API reference
│
├── src/                    # Custom components & pages
│   ├── components/         # React components
│   ├── css/                # Custom CSS
│   └── pages/              # Custom pages
│
├── static/                 # Static assets
│   ├── img/                # Images
│   └── .nojekyll           # GitHub Pages config
│
├── docusaurus.config.js    # Site configuration
├── sidebars.js             # Sidebar structure
└── package.json            # Dependencies
```

## ✍️ Writing Documentation

### All Examples Must Be Tested!

**Important**: Every code example in the documentation must be tested locally before adding it to the docs.

1. Add your example to `../docs-examples-test/main.go`
2. Run it: `cd ../docs-examples-test && go run main.go`
3. Capture the output
4. Add both code and output to documentation

### Example Template

Use this template for all examples:

```markdown
### Example: [Descriptive Title]

**Context**: Brief explanation of what we're doing.

**Code**:
```go
plygo.From(people).
    Where("Age").GreaterThan(30).
    Show()
```

**Output**:
```
╭──────────┬─────┬──────╮
│ Name     │ Age │ City │
├──────────┼─────┼──────┤
│ Alice    │  35 │ NYC  │
╰──────────┴─────┴──────╯
[1 rows × 3 columns]
```

**Notes**:
- 💡 Key insight
- ⚡ Performance tip
- 🎯 Best practice

:::tip Key Takeaway
Main point to remember
:::
```

### Code Blocks

Use proper language tags:

- `go` for Go code
- `bash` for shell commands
- `text` for plain output

### Admonitions (Callouts)

```markdown
:::tip Pro Tip
This is a tip
:::

:::info
This is information
:::

:::warning
This is a warning
:::

:::danger
This is dangerous
:::
```

## 🎨 Customization

### Theme Colors

Edit `src/css/custom.css` to change colors and styles.

### Sidebar

Edit `sidebars.js` to modify navigation structure.

### Config

Edit `docusaurus.config.js` for site-wide settings.

## 🚢 Deployment

### Automatic (GitHub Actions)

The site automatically deploys when you push to `main`:

```bash
git add .
git commit -m "docs: Update documentation"
git push origin main
```

GitHub Actions will:
1. Build the site
2. Deploy to GitHub Pages
3. Site available at: https://felipe-mansoldo.github.io/plyGO/

### Manual Deployment

```bash
npm run build
# Upload contents of build/ directory to your hosting
```

## 📝 Content Guidelines

1. **All examples tested**: Run every code snippet locally
2. **Show real output**: Use actual output from examples
3. **Clear context**: Explain what each example does
4. **Progressive complexity**: Start simple, add complexity gradually
5. **Consistent formatting**: Follow the example template
6. **Practical focus**: Real-world scenarios over academic examples

## 🛠️ Maintenance

### Adding New Pages

1. Create markdown file in appropriate directory
2. Add frontmatter with `sidebar_position`
3. Add to `sidebars.js` if needed
4. Test locally with `npm start`

### Updating Examples

1. Update code in `../docs-examples-test/`
2. Test: `go run main.go`
3. Capture output
4. Update documentation
5. Build: `npm run build`

## 📚 Resources

- [Docusaurus Documentation](https://docusaurus.io/)
- [Markdown Guide](https://docusaurus.io/docs/markdown-features)
- [Deployment Guide](https://docusaurus.io/docs/deployment)

## 🐛 Troubleshooting

### Build Fails

```bash
# Clean and rebuild
npm run clear
npm run build
```

### Broken Links

Check error message for specific broken link and fix the path.

### Port Already in Use

```bash
# Use different port
npm start -- --port 3001
```

## ✅ Pre-Deployment Checklist

- [ ] All examples tested locally
- [ ] All code blocks have correct language tags
- [ ] All links work
- [ ] Build succeeds: `npm run build`
- [ ] Serve works: `npm run serve`
- [ ] No console errors in browser
- [ ] Mobile responsive (test with browser dev tools)
- [ ] Search works (after deployment)

## 📧 Support

For issues or questions:
- GitHub Issues: https://github.com/felipe-mansoldo/plyGO/issues
- Documentation: https://felipe-mansoldo.github.io/plyGO/

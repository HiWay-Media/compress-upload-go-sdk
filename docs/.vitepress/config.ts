import { defineConfig } from 'vitepress'
//
// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: 'compress-upload-go-sdk',
  description: '',
  base: '/compress-upload-go-sdk/',
  themeConfig: {
    outline: [2, 3],
    socialLinks: [
      { icon: 'github', link: 'https://github.com/HiWay-Media/compress-upload-go-sdk' },
    ],
    sidebar: [
      { text: 'Introduction', link: '/' },
      { text: 'Getting started', link: '/getting-started' },
    ],
  },
})
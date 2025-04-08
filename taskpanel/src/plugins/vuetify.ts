/**
 * plugins/vuetify.ts
 *
 * Framework documentation: https://vuetifyjs.com`
 */

// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

// Composables
import { createVuetify } from 'vuetify'

const farmTheme = {
  dark: false,
  colors: {
    background: '#f5ddc3',
    'on-background': '#2a1a08',
    primary: '#285f8f',
    secondary: '#466860',
    accent: '#c45172',
    error: '#c45172',
    warning: '#673ab7',
    info: '#285f8f',
    success: '#466860'
  }
}

// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
export default createVuetify({
  theme: {
    defaultTheme: 'farmTheme',
    themes: {
      farmTheme,
    },
  },
})

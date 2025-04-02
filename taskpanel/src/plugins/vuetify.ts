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
    secondary: '#285f8f',
    accent: '#c45172',
    error: '#3f51b5',
    warning: '#673ab7',
    info: '#e91e63',
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

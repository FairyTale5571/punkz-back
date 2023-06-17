import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import dns from 'dns'

dns.setDefaultResultOrder('verbatim')

const PORT = process.env.PORT || 3000

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    // host: '0.0.0.0',
    port: PORT,
  },
  plugins: [react()],
})

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      backgroundImage: {
        'gradient-radial': 'radial-gradient(var(--tw-gradient-stops))',
        'body': "url('/img/bg.jpg')",
        'win-inactive': "url('/img/win_inactive.png')",
        'win-active': "url('/img/win_active.png')",
        'twitter-logo': "url('/img/twitter_logo.png')",
        'discord-logo': "url('/img/discord_logo.png')",
        'discord-button': "url('/img/discord_button.png')",
        'computer': "url('/img/computer.png')",
        'help': "url('/img/help.png')",
        'recycle': "url('/img/recycle.png')",
        'close': "url('/img/close.png')",
        'close-inactive': "url('/img/close_inactive.png')",
        'close-active': "url('/img/close_active.png')",
        'avatar': "url('/img/avatar.png')",
      },
      fontFamily: {
        segoe: "'Segoe UI 7'"
      },
      boxShadow: {
        hideButton: 'inset 0px 0px 0px 2px rgba(255, 255, 255, 0.25)'
      },
      dropShadow: {
        modal: '0px 4px 4px rgba(0, 0, 0, 0.25)'
      }
    }
  },
  plugins: [],
}
import { serve } from 'bun'

import { version } from '../package.json'

const routes = {
  '/api/version': () => Response.json({ version }, { status: 200 })
}

export default (port: number = 3000) => {
  return serve({
    port,
    routes,
    error(error) {
      console.error(error)
      return new Response('Internal Error', { status: 500 })
    }
  })
}
import Vue from 'vue'
import Router from 'vue-router'
import Meta from 'vue-meta'

import Landingpage01 from './views/landingpage01'
import NotFound from './views/not-found'
import './style.css'

Vue.use(Router)
Vue.use(Meta)
export default new Router({
  mode: 'history',
  routes: [
    {
      name: 'Landingpage01',
      path: '/',
      component: Landingpage01,
    },
    {
      name: '404 - Not Found',
      path: '**',
      component: NotFound,
      fallback: true,
    },
  ],
})

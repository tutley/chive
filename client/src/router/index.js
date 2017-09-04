import Vue from 'vue'
import Router from 'vue-router'
import Hello from '@/components/Hello'
import ExampleList from '@/components/examples/List'
import ExampleDetail from '@/components/examples/Detail'
import ExamplePost from '@/components/examples/Post'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Hello',
      component: Hello
    },
    {
      path: '/examples',
      name: 'Example List',
      component: ExampleList
    },
    {
      path: '/examples/:_id',
      name: 'Example Detail',
      component: ExampleDetail
    },
    {
      path: '/examples/post',
      name: 'Example Post',
      component: ExamplePost
    }
  ]
})

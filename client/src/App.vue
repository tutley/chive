<template>
<v-app id="foundirl">
    <v-navigation-drawer
      app
      fixed
      v-model="drawer"
      light
    >
      <v-list>
        <v-list-tile
          v-for="item in menuItems"
          :key="item.title"
          exact
          :to="item.link">
          <v-list-tile-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>{{ item.title }}</v-list-tile-content>
        </v-list-tile>
        <v-list-tile
          v-if="userIsAuthenticated"
          @click="onLogout">
          <v-list-tile-action>
            <v-icon>exit_to_app</v-icon>
          </v-list-tile-action>
          <v-list-tile-content>Logout</v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar app dark color="primary">
      <v-toolbar-side-icon @click.stop="drawer = !drawer"></v-toolbar-side-icon>
      <v-toolbar-title>
        Chive
      </v-toolbar-title>
     </v-toolbar>
      <v-content>
        <v-fab-transition>
          <v-btn fab color="accent" 
            bottom right fixed small
            @click="jump('/examples/post')">
            <v-icon dark>add</v-icon>
          </v-btn>
        </v-fab-transition>
        <router-view></router-view>
      </v-content>
  </v-app>
</template>

<script>
export default {
  name: 'app',
  data() {
    return {
      drawer: false
    }
  },
  computed: {
    userIsAuthenticated() {
      // This is where you would have some code to determine if the user is logged in
      return false
    },
    menuItems() {
      let menuItems = [
        { icon: 'home', title: 'Home', link: '/' },
        { icon: 'format_list_numbered', title: 'List Examples', link: '/examples' },
        { icon: 'add_circle', title: 'Post an Example', link: '/examples/post' }
        // { icon: 'lock_open', title: 'Sign in', link: '/signin' },
      ]
      if (this.userIsAuthenticated) {
        menuItems = [
          { icon: 'home', title: 'Home', link: '/' },
          { icon: 'format_list_numbered', title: 'List Examples', link: '/examples' },
          { icon: 'add_circle', title: 'Post an Example', link: '/examples/post' },
          { icon: 'account_circle', title: 'My Profile', link: '/profile' }
        ]
      }
      return menuItems
    }
  },
  methods: {
    jump(loc) {
      this.$router.push(loc)
    }
  },
  beforeMount() {
    // This removes the spinner once the app is loaded
    let pwadiv = document.getElementById('pwaloader')
    pwadiv.remove()
  }
}
</script>

<style lang="stylus">
$color-pack = false;

/* * local loading of material icons and roboto font */
@font-face {
  font-family: 'Material Icons';
  font-style: normal;
  font-weight: 400;
  src: url('https://fonts.gstatic.com/s/materialicons/v29/2fcrYFNaTjcS6g4U3t-Y5StnKWgpfO2iSkLzTz-AABg.ttf'); // fonts.gstatic.com/s/materialicons/v29/2fcrYFNaTjcS6g4U3t-Y5StnKWgpfO2iSkLzTz-AABg.ttf)
  format('truetype');
}

.material-icons {
  font-family: 'Material Icons';
  font-weight: normal;
  font-style: normal;
  font-size: 24px;
  line-height: 1;
  letter-spacing: normal;
  text-transform: none;
  display: inline-block;
  white-space: nowrap;
  word-wrap: normal;
  direction: ltr;
}

@font-face {
  font-family: 'Roboto';
  font-style: normal;
  font-weight: 300;
  src: local('Roboto Light'), local('Roboto-Light'), url('https://fonts.gstatic.com/s/roboto/v16/Hgo13k-tfSpn0qi1SFdUfaCWcynf_cDxXwCLxiixG1c.ttf');
  format('truetype');
}

@font-face {
  font-family: 'Roboto';
  font-style: normal;
  font-weight: 400;
  src: local('Roboto'), local('Roboto-Regular'), url('https://fonts.gstatic.com/s/roboto/v16/zN7GBFwfMP4uA6AR0HCoLQ.ttf') format('truetype');
}

@font-face {
  font-family: 'Roboto';
  font-style: normal;
  font-weight: 500;
  src: local('Roboto Medium'), local('Roboto-Medium'), url('https://fonts.gstatic.com/s/roboto/v16/RxZJdnzeo3R5zSexge8UUaCWcynf_cDxXwCLxiixG1c.ttf');
  format('truetype');
}

@font-face {
  font-family: 'Roboto';
  font-style: normal;
  font-weight: 700;
  src: local('Roboto Bold'), local('Roboto-Bold'), url('https://fonts.gstatic.com/s/roboto/v16/d-6IYplOFocCacKzxwXSOKCWcynf_cDxXwCLxiixG1c.ttf');
  format('truetype');
}
</style>

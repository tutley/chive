<template>
  <div class="mdl-grid">
    <div class="mdl-cell mdl-cell--8-col">
     <div class="demo-card-wide mdl-card mdl-shadow--2dp">
        <div class="mdl-card__title">
          <h2 class="mdl-card__title-text">{{ this.examples[$route.params._id].title }}</h2>
        </div>
        <div class="mdl-card__supporting-text">
          {{ this.examples[$route.params._id].body }}
        </div>
        <div class="mdl-card__actions mdl-card--border">
          <router-link class="mdl-button mdl-js-button mdl-button--raised mdl-button--colored" to="/post">
            ANSWER
          </router-link>
        </div>
        <div class="mdl-card__menu">
          <button class="mdl-button mdl-button--icon mdl-js-button mdl-js-ripple-effect">
            <i class="material-icons">share</i>
          </button>
        </div>
      </div>

    </div>
  </div>
</template>

<script>
import data from '../../data'
export default {
  data () {
    return {
      'examples': data.examples
    }
  }
}
</script>
<style scoped>
.demo-card-wide.mdl-card {
  width: 512px;
}
.demo-card-wide > .mdl-card__title {
  color: #fff;
  height: 176px;
}
.demo-card-wide > .mdl-card__menu {
  color: #fff;
}

  .actions {
    text-align: center;
  }
</style>

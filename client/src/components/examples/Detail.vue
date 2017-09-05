<template>
  <div class="mdl-grid">
    <div class="mdl-cell mdl-cell--8-col">
     <div class="demo-card-wide mdl-card mdl-shadow--2dp">
        <div v-show="!this.editable" class="mdl-card__title">
          <h2 class="mdl-card__title-text">{{ this.example.title }}</h2>
        </div>
        <div v-show="!this.editable" class="mdl-card__supporting-text">
          {{ this.example.body }}
        </div>
        <form>
        <div v-show="this.editable" class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label is-upgraded is-dirty">
          <input id="title" v-model="example.title" type="text" class="mdl-textfield__input"/>
          <label for="title" class="mdl-textfield__label">Title</label>
        </div>
        <div v-show="this.editable" class="mdl-textfield mdl-js-textfield">
          <textarea class="mdl-textfield__input" type="text" rows= "3" v-model="example.body" id="body" ></textarea>
        </div>
        <div class="mdl-card__actions mdl-card--border">
          <div v-show="!this.editable && !this.delete">
            <a class="mdl-button mdl-js-button mdl-button--raised mdl-button--colored" @click.prevent="doEdit(example)">
              EDIT
            </a>
            <span class="fill-space">&nbsp;</span>
            <a class="mdl-button mdl-js-button mdl-button--raised mdl-button--accent" @click.prevent="doDelete()">
              DELETE
            </a>
          </div>
          <div v-show="this.editable">
            <a class="mdl-button mdl-js-button mdl-button--raised mdl-button--colored" @click.prevent="saveUpdate()">
              SAVE
            </a>
            <a class="mdl-button mdl-js-button mdl-button--raised mdl-button--colored" @click.prevent="cancelEdit()">
              CANCEL
            </a>
          </div>
          <div v-show="this.delete">
          <div>
            <span class="mdl-chip mdl-chip--basic">
              <span class="mdl-chip__text">Are you sure? Think of the children!</span>
            </span>
            </div>
            <a class="mdl-button mdl-js-button mdl-button--raised mdl-button--accent" @click.prevent="sendDelete()">
              DELETE
            </a>
            <a class="mdl-button mdl-js-button mdl-button--raised mdl-button--colored" @click.prevent="cancelDelete()">
              CANCEL
            </a>
          </div>
        </div>
        </form>
      </div>
      <ul v-if="this.errors && this.errors.length">
        <li v-for="error of this.errors">
          <span>{{error.message}}</span>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  import {HTTP} from '../../api'

  export default {
    methods: {
      doEdit: function (example) {
        this._originalExample = Object.assign({}, example)
        this.editable = true
      },
      cancelEdit: function () {
        Object.assign(this.example, this._originalExample)
        this.editable = false
      },
      saveUpdate: function () {
        HTTP.put('examples/' + this.$route.params.id, this.example)
        .then(response => {
          this.editable = false
        })
        .catch(e => {
          this.errors.push(e)
        })
      },
      doDelete: function () {
        this.delete = true
      },
      cancelDelete: function () {
        this.delete = false
      },
      sendDelete: function () {
        HTTP.delete('examples/' + this.$route.params.id)
        .then(response => {
          this.$router.push({name: 'Example List'})
        })
        .catch(e => {
          this.errors.push(e)
        })
      }
    },
    data: () => ({
      example: {},
      errors: [],
      editable: false,
      delete: false,
      _originalExample: {}
    }),
    created () {
      HTTP.get('examples/' + this.$route.params.id)
      .then(response => {
        this.example = response.data
      })
      .catch(e => {
        this.errors.push(e)
      })
    }
  }
</script>
<style scoped>
.fill-space {
  // This fills the remaining space, by using flexbox.
  // Every toolbar row uses a flexbox row layout.
  flex: 1 1 auto;
}
.demo-card-wide.mdl-card {
  width: 512px;
}
.demo-card-wide > .mdl-card__title {
  height: 176px;
}
.demo-card-wide > .mdl-card__menu {
  color: #fff;
}

  .actions {
    text-align: center;
  }
</style>

<template>
  <v-layout
    row
    class="align-center layout app--page-header"
  >
    <div class="page-header-left">
      <h3 class="pr-3">
        {{ title }}
      </h3>
    </div>
    <v-breadcrumbs
      :items="breadcrumbs"
      divider=">"
    />
    <v-spacer />
  </v-layout>
</template>

<script>
import menu from '@/api/menu'

export default {
    data () {
        return {
            title: ''
        }
    },
    computed: {
        breadcrumbs: function () {
            let breadcrumbs = [
                {
                    disabled: false,
                    icon: 'home',
                    href: '/'
                }
            ]
            menu.forEach((item) => {
                if (item.items) {
                    let child = item.items.find(i => {
                        return i.component === this.$route.name
                    })

                    if (child) {
                        let p = {
                            text: this.$t('menu.' + item.title),
                            disabled: false,
                            href: item.path
                        }

                        let c = {
                            text: this.$t('menu.' + child.title),
                            disabled: false,
                            href: child.path
                        }

                        breadcrumbs.push(p)
                        breadcrumbs.push(c)
                    }
                } else {
                    if (item.name === this.$route.name) {
                        let p = {
                            text: this.$t('menu.' + item.title),
                            disabled: false,
                            href: item.path
                        }

                        // this.title = item.title;
                        breadcrumbs.push(p)
                    }
                }
            })
            return breadcrumbs
        }
    }
}
</script>

<template>
  <ControlApp
      guide_url="/#"
      :customMenuItems="customMenuItems"
  >

    <template v-slot:headerMenu>

    </template>
    <template v-slot:sideBottomMenu>
      <li v-for="customMenuItem in customMenuItems" :key="customMenuItem.index">
        <router-link :to="customMenuItem.link">
          <span :class="customMenuItem.icon"></span>
          <span>{{ customMenuItem.title }} {{ $user.role }} {{ $user.login }}</span>

        </router-link>
      </li>
    </template>

  </ControlApp>
</template>

<script>
import ControlApp from "../controlTemplate/app"
import {mapGetters} from "vuex";

export default {
  computed: {

    ...mapGetters({

      themeMode: "themeMode",
    }),
    customMenuItems() {
        return [];
    }
  },
  methods: {

  },
  components: {
    ControlApp
  },
  mounted() {

  },
  data() {
    return {

    }
  },
}
</script>

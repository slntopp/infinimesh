<template>
  <a-dropdown :trigger="['contextmenu']">
    <div
      @click.meta.exact="$emit((selected ? 'de' : '') + 'select', device.id)"
    >
      <a-card
        :hoverable="true"
        :bordered="false"
        :ref="`device-card-${device.id}`"
        :class="selected ? 'card-selected' : ''"
        @click.left.exact="
          $router.push({
            name: 'dashboard-devices-id',
            params: { id: device.id },
          })
        "
      >
        <template slot="title">{{ device.name }}</template>
        <template slot="extra">
          <b class="muted">{{ device.id }}</b>
          <a-tooltip
            :title="device.enabled ? 'Device enabled' : 'Device is not enabled'"
            placement="bottom"
          >
            <a-icon
              type="bulb"
              :style="{ color: device.enabled ? '#52c41a' : '#eb2f96' }"
              theme="filled"
            />
          </a-tooltip>
        </template>
        <template>
          <a-row v-if="device.tags.length">
            <a-col :span="4" :xl="3">Tags:</a-col>
            <a-col :span="20" :xl="21">
              <a-tag
                v-for="tag in device.tags"
                :key="tag"
                style="margin-bottom: 5px"
                @click.left.exact="
                  (e) => {
                    e.stopPropagation();
                    $emit('tag-clicked', tag);
                  }
                "
                >{{ tag }}
              </a-tag>
            </a-col>
          </a-row>
          <a-row v-else type="flex" justify="center" class="muted"
            >No tags were provided</a-row
          >
        </template>
        <template v-if="state">
          <device-state
            :title="false"
            :version="false"
            :state="state.reported"
          />
        </template>
      </a-card>
    </div>
    <a-menu slot="overlay">
      <a-menu-item key="open">
        <nuxt-link
          :to="{ name: 'dashboard-devices-id', params: { id: device.id } }"
        >
          <a-button type="link"> Open </a-button>
        </nuxt-link>
      </a-menu-item>
      <a-menu-item key="toggle">
        <a-button type="link" @click="handleToggleDevice(false)">
          {{ device.enabled ? "Disable" : "Enable" }}
        </a-button>
      </a-menu-item>
      <a-menu-item key="toggle-selection">
        <a-button
          type="link"
          @click="$emit((selected ? 'de' : '') + 'select', device.id)"
        >
          {{ selected ? "Deselect" : "Select" }}
        </a-button>
      </a-menu-item>
      <a-menu-item key="select-all">
        <a-button type="link" @click="$emit('select-all')">
          Select All
        </a-button>
      </a-menu-item>
    </a-menu>
  </a-dropdown>
</template>

<script>
import Vue from "vue";
import deviceControlMixin from "@/mixins/device-control";
import DeviceState from "@/components/device/State";

export default Vue.component("device-list-card", {
  mixins: [deviceControlMixin],
  components: { DeviceState },
  props: {
    device: {
      required: true,
      type: Object,
    },
    selected: {
      required: false,
      default: false,
      type: Boolean,
    },
  },
  computed: {
    state() {
      let state = this.$store.getters["devices/get_state"](this.device.id);
      console.log(state, this.device.id);
      if (!state) {
        return false;
      }
      return state;
    },
  },
});
</script>

<style scoped>
.card-selected {
  --card-selected-shadow: 20px 15px 10px 5px rgba(0, 0, 0, 0.7);
  -webkit-box-shadow: var(--card-selected-shadow);
  -moz-box-shadow: var(--card-selected-shadow);
  box-shadow: var(--card-selected-shadow);
}
</style>
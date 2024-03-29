<template>
  <div id="devicesTable">
    <a-row
      type="flex"
      style="padding: 4px 10px"
      class="tile-bar tile-bar-transparent"
      align="middle"
    >
      <template>
        <a-col>
          <a-input
            placeholder="Search device..."
            class="devices-search-input"
            :allowClear="true"
            v-model="searchQuery"
          >
            <a-select
              slot="addonBefore"
              class="device-search-prefix"
              v-model="searchKey"
            >
              <a-select-option value="all"> Everywhere </a-select-option>
              <a-select-option value="name"> Names </a-select-option>
              <a-select-option value="id"> IDs </a-select-option>
              <a-select-option value="tags"> Tags </a-select-option>
              <a-select-option value="namespace"> Namespace </a-select-option>
            </a-select>
          </a-input>
        </a-col>
        <a-col>
          <a-row type="flex" justify="center">
            <a-col>
              <a-switch
                id="group-by-tags-switch"
                un-checked-children="Group by tags"
                checked-children="Whole registry"
                v-model="groupByTags"
              />
            </a-col>
          </a-row>
        </a-col>
      </template>

      <template v-if="selectedDevices.length">
        <a-col>
          <a-button type="link" @click="selectedDevices = []" icon="close"
            >Deselect all
          </a-button>
        </a-col>
        <a-col>
          <a-button
            type="success"
            style="margin-right: 5px"
            @click="toggleSelected(true)"
            >Enable All
          </a-button>
          <a-button type="danger" @click="toggleSelected(false)"
            >Disable All
          </a-button>
        </a-col>
      </template>
      <template v-else>
        <a-col>
          <a-button
            type="success"
            @click="selectedDevices = suggested.map((d) => d.id)"
            >Select All
          </a-button>
        </a-col>
      </template>
      <template v-if="pool.length">
        <a-col>
          <a-button type="success" @click="handleLoadState"
            >Load Devices State</a-button
          >
        </a-col>
      </template>
    </a-row>
    <template v-if="groupByTags">
      <a-collapse
        :bordered="false"
        accordion
        class="tags-collapse-tiles-wrap"
        :activeKey="tagsCollapseActive"
      >
        <a-collapse-panel
          v-for="tag in tags"
          :key="tag"
          class="tags-collapse-tile-wrap"
        >
          <span
            slot="header"
            style="color: var(--text-color); font-size: 18px"
            >{{ tag }}</span
          >
          <div slot="extra" @click="(e) => e.stopPropagation()">
            <a-button
              type="success"
              style="border-radius: 100px; height: 24px"
              @click="
                selectedDevices.push(
                  ...suggested
                    .filter((d) => d.tags.includes(tag))
                    .map((d) => d.id)
                )
              "
              >Select All
            </a-button>
          </div>

          <device-pool
            :div="div"
            :selected="selectedDevices"
            :pool="suggested.filter((d) => d.tags.includes(tag))"
            :grouped="true"
            @select="(id) => selectedDevices.push(id)"
            @deselect="
              (id) => selectedDevices.splice(selectedDevices.indexOf(id), 1)
            "
            @select-all="selectedDevices = suggested.map((d) => d.id)"
            style="
              background-color: var(--background-color);
              border-radius: var(--border-radius-base);
              padding-bottom: 10px;
            "
          />
        </a-collapse-panel>
      </a-collapse>
    </template>
    <device-pool
      :div="div"
      :selected="selectedDevices"
      :pool="mainPool"
      @select="(id) => selectedDevices.push(id)"
      @deselect="(id) => selectedDevices.splice(selectedDevices.indexOf(id), 1)"
      @select-all="selectedDevices = suggested.map((d) => d.id)"
      @tag-clicked="
        (tag) => {
          groupByTags = true;
          tagsCollapseActive = tag;
        }
      "
    >
      <a-row
        class="create-form"
        type="flex"
        justify="center"
        align="middle"
        slot="device-create-form"
      >
        <nuxt-link
          v-if="user.default_namespace.id === namespace"
          :to="{ name: 'dashboard-namespaces', query: { create: true } }"
          no-prefetch
        >
          <h3 style="padding: 15px">
            <p>
              You can't create devices in your root namespace, switch to another
              one to perform device create.
            </p>
            <p>
              Click here to create new namespace, or switch namespace on top of
              the page.
            </p>
          </h3>
        </nuxt-link>
        <template v-else>
          <a-icon
            type="plus"
            style="font-size: 6rem; height: 100%; width: 100%"
            @click="addDeviceActive = true"
          />
          <device-add
            :active="addDeviceActive"
            @cancel="addDeviceActive = false"
            @add="handleDeviceAdd"
          />
        </template>
      </a-row>
    </device-pool>
  </div>
</template>

<script>
import DevicePool from "@/components/device/Pool.vue";
import DeviceAdd from "@/components/device/Add.vue";

const divs = {
  xs: 1,
  sm: 2,
  md: 2,
  lg: 3,
  xl: 3,
  xxl: 4,
};

export default {
  name: "devicesTable",
  components: {
    DeviceAdd,
    DevicePool,
  },
  data() {
    return {
      addDeviceActive: false,
      selectedDevices: [],
      groupByTags: false,
      tagsCollapseActive: "",
      searchKey: "all",
      searchQuery: "",
    };
  },
  watch: {
    selectedDevices() {
      this.selectedDevices.filter((e, i, self) => self.indexOf(e) === i);
    },
    searchQuery() {
      let lower = this.searchQuery.toLowerCase();
      if (lower !== this.searchQuery) this.searchQuery = lower;

      if (this.searchQuery.includes(":")) {
        let new_key;
        let colon = this.searchQuery.indexOf(":");
        if (
          ["all", "name", "id", "tags", "namespace"].includes(
            (new_key = this.searchQuery.slice(0, colon))
          )
        )
          this.searchKey = new_key;
        this.searchQuery = this.searchQuery.slice(colon + 1);
      }
    },
  },
  computed: {
    tags() {
      return this.pool.reduce((r, el) => {
        el.tags.forEach((t) => r.add(t));
        return r;
      }, new Set());
    },
    user() {
      if (!this.$store.getters.loggedInUser) return { default_namespace: {} };
      return this.$store.getters.loggedInUser;
    },
    namespace() {
      return this.$store.getters["devices/currentNamespace"];
    },
    pool: {
      deep: true,
      get() {
        return this.$store.state.devices.pool;
      },
    },
    suggested: {
      deep: true,
      get() {
        if (this.searchQuery === "") {
          return this.pool;
        }
        let searchable;
        if (this.searchKey == "all") {
          searchable = (device) => {
            return [
              device.id,
              device.name,
              device.namespace,
              device.tags.join(","),
            ]
              .join("|")
              .toLowerCase();
          };
        } else {
          searchable = (device) => device[this.searchKey].toLowerCase();
        }
        return this.pool
          .map((d) => {
            d.searchable = searchable(d);
            return d;
          })
          .filter((d) => d.searchable.indexOf(this.searchQuery) !== -1)
          .sort((d) => -d.searchable.indexOf(this.searchQuery));
      },
    },
    mainPool: {
      deep: true,
      get() {
        return [{ type: "create-form" }, ...this.suggested];
      },
    },
    div() {
      return divs[this.$store.state.window.gridSize];
    },
  },
  mounted() {
    this.$store.commit("window/setTopAction", {
      icon: "undo",
      callback: this.refresh,
    });
  },
  beforeDestroy() {
    this.$store.commit("window/unsetTopAction");
  },
  methods: {
    refresh() {
      this.$store.dispatch("devices/get");
    },
    handleLoadState() {
      this.$store.dispatch("devices/state");
    },
    handleDeviceAdd(device) {
      this.$store.dispatch("devices/add", {
        device: device,
        error: (err) => {
          this.$notification.error({
            message: "Failed to create the device",
            description: `Response: ${err.response.data.message}`,
            duration: 10,
          });
        },
        always: () => {
          this.addDeviceActive = false;
        },
      });
    },
    toggleSelected(enable) {
      let vm = this;
      this.updateSelected(
        () => {
          return {
            enabled: enable,
          };
        },
        {
          success: (response) => {
            let res = response.reduce(
              (r, el) => {
                r[el.status == 200 ? "success" : "fail"]++;
                return r;
              },
              {
                success: 0,
                fail: 0,
              }
            );
            vm.$notification.info({
              message: `${enable ? "Enabled" : "Disabled"} ${
                response.length
              } devices`,
              description: `Result: success - ${res.success}, failed: ${res.fail}.`,
            });
          },
        }
      );
    },
    updateSelected(modifier, { success, error }) {
      let patchPromises = this.pool
        .filter((d) => this.selectedDevices.includes(d.id))
        .map((device) => {
          return this.$axios({
            url: `/api/devices/${device.id}`,
            method: "patch",
            data: modifier(device),
          });
        });
      Promise.all(patchPromises)
        .then((res) => {
          if (success) success(res);
        })
        .catch((err) => {
          if (error) error(err);
        })
        .then(() => {
          this.$store.dispatch("devices/get");
        });
    },
  },
};
</script>

<style scoped>
.create-form {
  border-radius: var(--border-radius-base);
  background: var(--primary-color)-dark;
  border: var(--border-base);
  min-height: 8rem;
  cursor: pointer;
}
.create-form .anticon {
  color: var(--icon-color-dark);
}
.tile-bar {
  margin-top: 10px;
  width: 100%;
  background: var(--primary-color);
  border-radius: var(--border-radius-base);
  color: var(--line-color);
  min-height: 32px;
}
.tile-bar-transparent {
  background: none;
}
.tile-bar .ant-btn-link {
  color: var(--line-color) !important;
}
.tile-bar .ant-btn {
  height: 90%;
  border-radius: 100px;
}
#group-by-tags-switch {
  background-color: var(--switch-color);
}
#group-by-tags-switch.ant-switch-checked {
  background-color: var(--success-color);
}
.devices-search-input {
  height: 90%;
  border: 1px solid var(--background-color);
  min-width: 256px;
  border-radius: 100px;
}
@media (max-width: 768px) {
  .tile-bar > [class*="ant-col"] {
    margin-top: 3px;
    margin-bottom: 3px;
  }
}
.tile-bar > .ant-col + .ant-col {
  margin-left: 15px;
}
.tags-collapse-tiles-wrap {
  margin-top: 10px;
  margin-bottom: 5px;
}
.tags-collapse-tile-wrap {
  border-radius: var(--border-radius-base);
  background-color: var(--primary-color);
}
.tags-collapse-tile-wrap:last-child {
  border-radius: var(--border-radius-base);
}
.ant-collapse-borderless {
  background-color: var(--background-color) !important;
}
.ant-collapse-borderless > .ant-collapse-item {
  border-bottom: 1px solid var(--background-color);
}
</style>
<style>
.devices-search-input
  > .ant-input-group
  > .ant-input-affix-wrapper:not(:first-child)
  .ant-input {
  border-radius: 0 100px 100px 0 !important;
}
.device-search-prefix {
  height: 90%;
}
.devices-search-input > .ant-input-group > .ant-input:first-child,
.ant-input-group-addon:first-child {
  border-radius: 100px 0 0 100px !important;
}
</style>
<template>
  <div>
    <el-row class="title">
      <h1>Modbus Config</h1>
    </el-row>
    <el-row class="title">
      <el-image style="width:30%;" :src="url">
      </el-image>
    </el-row>
    <el-row>
      <el-form :model="dynamicValidateForm" ref="dynamicValidateForm" label-width="200px" class="demo-dynamic">
        <el-row v-for="(di, index) in dynamicValidateForm.dis" :key="di.key">
          <el-col :span=2 :offset=2>
            <div class="label-di">
              <div class="dot" v-bind:class="{ active: di.status, unactive: !di.status}"></div>
              <span>{{'DI-' + (index+1)}}</span>
              <el-button size="mini" type="danger" @click.prevent="remove(di) ">{{$t('common.Remove')}}
              </el-button>
            </div>
          </el-col>
          <el-col :span=14>
            <el-form-item :rules="httpRules[0]" :prop="`dis[${index}].high`" :key="'high' + di.name">
              <label slot="label">
                {{$t('modbus.rising_edge_triggered')}}
                <el-tooltip placement="bottom">
                  <div slot="content">{{$t('modbus.high_tip')}}</div>
                  <span><i class="el-icon-question"></i>:</span>
                </el-tooltip>
              </label>
              <el-input size="mini" v-model="di.high" placeholder="http:// or https:// ">
                <template #append>
                  <el-button icon="el-icon-s-promotion" @click="tryUrl(di.high)"></el-button>
                </template>
              </el-input>
            </el-form-item>

            <el-form-item :rules="httpRules[0]" :prop="`dis[${index}].low`" :key="'low' + di.name">
              <label slot="label">
                {{$t('modbus.trailing_edge_triggered')}}
                <el-tooltip placement="bottom">
                  <div slot="content">{{$t('modbus.low_tip')}}</div>
                  <span><i class="el-icon-question"></i>:</span>
                </el-tooltip>
              </label>
              <el-input size="mini" v-model="di.low" placeholder="http:// or https:// ">
                <template #append>
                  <el-button icon="el-icon-s-promotion" @click="tryUrl(di.low)"></el-button>
                </template>
              </el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row class="submit">
          <el-form-item>
            <el-button type="primary" @click="submit('dynamicValidateForm')">{{$t('common.Submit')}}</el-button>
            <el-button @click="add">{{$t('modbus.new_di')}}</el-button>
          </el-form-item>
        </el-row>
      </el-form>
    </el-row>
  </div>
</template>

<style scoped>

.el-form-item__content .el-input-group {
    vertical-align: inherit;
}

/* title 居中加大 */
.title {
    display: flex;
    justify-content: center;
    font-size: large;
}

/*label-di */
.label-di{
    margin-top: 10px;
    height: 85px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-evenly;
}

/* dot DI 状态 */
.dot {
    height: 14px;
    width: 14px;
    border-radius: 50%;
    display: inline-block;
  }
.active {
    background-color: red;
}
.unactive {
    background-color: gray;
}

.submit {
  display: flex;
  justify-content: center;
}
</style>

<script>
module.exports = {
  data() {
    return {
      httpRules: [{validator: this.validateURL, message: '请输入正确 HTTP 服务器地址', trigger: ['blur'] }],
      dynamicValidateForm: {
        dis: [{
          name: "1",
          status: false,
          low: "",
          high: ""},
        ],
      },

      url: "./assists/device.jpg",
      ws: null,

      loader: null,
      main: null,
    };
  },
  components: {},
  mounted() {
    this.getConfig();
    this.openWs();
  },
  methods: {
    validateURL(rule, value, callback) {
        if(!value){
          callback()
          return 
        }
        if(value.startsWith("http")){
          callback()
          return
        }
        callback(rule.message);
        return
    },
    getConfig() {
      fetch(`http://${window.location.host}/modbus/config`)
        .then((res) => {
          return res.json();
        })
        .then((res) => {
          if(res.data){
            this.dynamicValidateForm.dis = res.data.dis.map((item, index) => {
              item.status = false;
              item.name = `${index +1}`;
              return item;
            });
          }
        });
    },
    openWs() {
      this.ws = new WebSocket(`ws://${window.location.host}/modbus/status`);
      this.ws.onopen = function (evt) {
        console.log("OPEN");
      };
      this.ws.onclose = function (evt) {
        console.log("CLOSE");
        this.ws = null;
      };
      this.ws.onmessage = (evt) => {
        console.log("RESPONSE: " + evt.data);
        evt.data.split("").forEach((item, index) => {
         let  data  = this.dynamicValidateForm.dis[index]
          if(data){
            data.status = item == "1";
          }
        });
      };
      this.ws.onerror = function (evt) {
        console.log("ERROR: " + evt.data);
      };
    },
    submit(form) {
      this.$refs[form].validate(async (valid) => {
        if (!valid){
          return
        }
        fetch(`http://${window.location.host}/modbus/config`, {
          body: JSON.stringify(this.dynamicValidateForm),
          method: "POST",
        })
          .then((res) => {
            if(res.ok){
              this.$message.success(this.$t("common.OperationSuccessful"))
            }else{
              this.$message.error(this.$t("common.OperationFailed"))
            }
          })
          .catch((err) => {
            console.error(err);
          });
      })

    },
    //删除 DI
    remove(item) {
      var index = this.dynamicValidateForm.dis.indexOf(item);
      if (index !== -1) {
        this.dynamicValidateForm.dis.splice(index, 1);
      }
    },
    // 添加 DI
    add() {
      this.dynamicValidateForm.dis.push({
          name: `${this.dynamicValidateForm.dis.length + 1}`,
          status: false,
          low: "",
          high: "",
      });
    },
    // 测试触发
    tryUrl(url){
      if(!url){
        return
      }
      fetch(`http://${window.location.host}/modbus/play-mp3`, {
        body: JSON.stringify({url}),
        method: 'POST'
      })
        .then((res) => {
          if(res.ok){
            this.$message.success(this.$t("common.OperationSuccessful"))
          }else{
            this.$message.error(this.$t("common.OperationFailed"))
          }
        })
        .catch(() => {
          this.$message(this.$t("common.OperationFailed"))
        })
    }
  },
};

</script>

<template>
  <div id="new-article">
    <input v-model="title" class="title" placeholder="タイトル(Title)" type="text">
    <div class="lines" v-for="n in num" :key="n">
      <template v-if="!linkBool[(n-1)]">
        <textarea maxlength="255" :class="boldClass[(n-1)]" v-model="lines[(n-1)]" ref="textarea" :style="changeHeight(n)" @keydown="methodsHappeningByTyped(n)" type="text" placeholder="New Line" wrap=""></textarea>
      </template>
      <template v-else>
        <a target="blank" class="link" :href="lines[(n-1)]">{{lines[(n-1)]}}</a>
      </template>
      <template v-if="restNumbers[(n-1)] >= 0">
        <small class="rest-numbers">{{restNumbers[(n-1)]}}</small>
      </template>
      <template v-else>
        <span class="over-numbers">Over the number of characters</span>
      </template>
      <i :style="linkStyle[(n-1)]" @click="changeToLink(n)" class="material-icons grey-text lighten-1 type">attach_file</i>
      <span :style="boldStyle[(n-1)]" @click="getBold(n)" class="bold grey-text lighten-1 type">B</span>
    </div>
    <button @click="newLine()" class="add-btn grey lighten-5"><i class="grey-text lighten-1 material-icons">add</i></button>
    <button class="btn grey share-btn" @click="createArticle()">Share</button>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      title: "",
      lines: [""],
      linkNum: [],
      boldNum: [],
      num: 1,
      height: [40],
      restNumbers: [254],
      boldClass: [""],
      boldBool: false,
      boldStyle: [
        {"color": ""},
      ],
      linkBool: [false],
      linkStyle: [
        {"color": ""}
      ],
    }
  },
  computed: {
    changeHeight() {
      self = this;
      return function(n) {
        return {
          "height": self.height[n-1] + "px",
        }
      }
    }
  },
  watch: {
  },
  methods: {
    newLine() {
      this.num ++;
      this.lines.push("");
      this.height.push(40);
      this.restNumbers.push(254);
      this.boldClass.push("");
      this.boldStyle.push({"color": ""});
      this.linkBool.push(false);
      this.linkStyle.push({"color": ""});
    },
    methodsHappeningByTyped(n) {
      this.adjustHeight(n);
      this.getRestNumbers(n);
    },
    adjustHeight(n){
      var textareaNum = n - 1 - this.linkNum.length;
      var linage = this.$refs.textarea[textareaNum].value.split('\n').length + 1;
      const RealHeight = this.$refs.textarea[textareaNum].scrollHeight
      this.height.splice((n-1), 1, (linage * 16 + 10));
    },
    getRestNumbers(n) {
      var textareaNum = n - 1 - this.linkNum.length;
      var restNum = (254 - this.$refs.textarea[textareaNum].value.length)
      this.restNumbers.splice((n-1), 1, restNum)
    },
    getBold(n) {
      if (this.boldBool === false) {
        this.boldClass.splice((n-1), 1, "getBold");
        this.boldBool = true;
        this.boldStyle.splice((n-1), 1, {"color": "#333 !important"});
        this.boldNum.push(n);
      } else {
        this.boldClass.splice((n-1), 1, "");
        this.boldBool = false;
        this.boldStyle.splice((n-1), 1, {"color": ""});
        for(let i = 0; i < this.boldNum.length; i ++) {
          if(this.boldNum[i] === n) {
            this.boldNum.splice(i, 1)
          }
        };
      }
    },
    changeToLink(n) {
      if (this.linkBool[(n-1)] === false) {
        this.linkBool.splice((n-1), 1, true);
        this.linkStyle.splice((n-1), 1, {"color": "#64b5f6 !important"});
        this.linkNum.push(n);
      } else {
        this.linkBool.splice((n-1), 1, false);
        this.linkStyle.splice((n-1), 1, {"color": ""});
        for(let i = 0; i < this.linkNum.length; i ++) {
          if(this.linkNum[i] === n) {
            this.linkNum.splice(i, 1)
          }
        };
      }
    },
    createArticle() {
      for (let i = 0; i < (this.lines.length -1); i ++) {
        this.lines[i] += "&divide;"
      }
      const params = new URLSearchParams();
      params.append("title", this.title);
      params.append("lines", this.lines);
      console.log(this.lines)
      params.append("link-numbers", this.linkNum);
      params.append("bold-numbers", this.boldNum);
      axios.post("/api/articles", params).then((res) => {
        this.$router.push(`/article/${res.data.data.id}`);
      })
    }
  },
}
</script>


<style lang="scss" scoped>
#new-article {
  width: 600px;
  margin: auto;
  margin-bottom: 80px;

  .title {
    border: none;
    font-size: 1.5rem;
    font-weight: bold;
    margin-bottom: 80px;
  }
  .title:focus {
    outline: none;
  }
  .lines {
    display: flex;
    margin-bottom: 10px;
    position: relative;
    textarea {
      border: none;
      white-space: nowrap;
      resize: none;
    }
    textarea:focus {
      outline: none;
    }
    .type {
      cursor: pointer;
    }
    .link {
      width: 564px;
      height: 42px;
    }
    .rest-numbers {
      color: #999;
      position: absolute;
      right: 0;
      bottom: -5px;
    }
    .over-numbers {
      color: tomato;
      position: absolute;
      right: 0;
      bottom: -5px;
    }
    .bold {
      font-weight: bold;
      font-size: 1.25rem;
    }
    .getBold {
      font-weight: bold;
    }
  }
  .add-btn {
    width: 50px;
    height: 50px;
    border: 1px solid #e0e0e0;
    border-radius: 50%;
  }
  .share-btn {
    position: fixed;
    right: 50px;
    bottom: 30px;
  }
}
@media (max-width: 480px) {
  #new-article {
    width: 98%;
    margin: auto;
    margin-bottom: 100px;
    .share-btn {
      bottom: 60px;
    }
  }
}
</style>
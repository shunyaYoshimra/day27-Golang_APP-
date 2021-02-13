<template>
  <div id="new-posts">
    <div class="contents">
      <p>留学をシェアする(Share your study abroad)</p>
      <div class="error">{{err}}</div>
      <form ref="myForm" :action="getPath" method="post" enctype="multipart/form-data">
        <textarea name="content" placeholder="内容(Content)"></textarea>
        <div class="tags">
          <span v-for="(eachTag, id) in tags" :key="id">#{{eachTag}}</span>
        </div>
        <hr />
        <input name="file" type="file" id="post-file" class="hide" multiple="multiple"/>
        <label for="post-file">
          <i class="material-icons">image</i>
        </label>
        <a href="#modal1" class="modal-trigger">
          <i class="material-icons">local_offer</i>
        </a>
        <br />
        <input type="hidden" name="time" :value="now" />
        <input value="share" type="submit" class="btn grey lighten-1" />
        <div id="modal1" class="modal">
          <div class="modal-content">
            <p>関連タグ(Tags)</p>
            <small>留学した国名、大学名など関連するタグをつけてみましょう。タグは複数個つけれます。</small>
            <br />
            <div class="addTags">
              <span v-for="(eachTag, id) in tags" :key="id">
                #{{eachTag}}
                <i
                  @click="removeTags(eachTag)"
                  class="material-icons tiny black-text removeTag"
                >close</i>
              </span>
            </div>
            <input type="text" name="tags" v-model="tag" />
          </div>
          <div class="modal-footer">
            <span @click="addTags()" class="btn-floating btn-small waves-effect blue lighten-3">
              <i class="material-icons white-text">add</i>
            </span>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      err: "",
      tag: "",
      tags: ["tag"],
      now: "00:00:00:00:00:000"
    };
  },
  mounted() {
    // jquery for modal
    $(document).ready(function() {
      $(".modal").modal();
    });
    // get now
    const date = new Date();
    this.now =
      date.getMonth() +
      ":" +
      date.getDate() +
      ":" +
      date.getHours() +
      ":" +
      date.getMinutes() +
      ":" +
      date.getSeconds() +
      ":" +
      date.getMilliseconds();
  },
  computed: {
    getPath: function() {
      return "/api/posts/" + this.tags;
    }
  },
  methods: {
    addTags() {
      this.tags.push(this.tag);
      this.tag = "";
    },
    removeTags(val) {
      var tag = val;
      var idx = this.tags.indexOf(tag);
      if (idx >= 0) {
        this.tags.splice(idx, 1);
      }
    }
  }
};
</script>

<style lang="scss" scoped>
#new-posts {
  width: 100%;
  height: 650px;
  background-color: #fff;

  .contents {
    margin: 0 auto;
    padding-top: 30px;
    width: 550px;

    .error {
      text-align: center;
      color: tomato;
    }

    p {
      font-weight: bold;
    }

    form {
      textarea {
        background-color: #fafafa;
        height: 300px;
        border: 1px solid #333;
        font-size: 1rem;
        padding: 1rem 1rem;
      }

      .tags {
        span {
          margin-right: 1rem;
          color: #64b5f6;
        }
      }

      textarea:focus {
        outline: none;
      }

      i {
        color: #64b5f6;
      }
    }
    .modal {
      p {
        font-weight: bold;
      }
      .addTags {
        display: flex;
        flex-wrap: wrap;
        margin: 1rem 0;

        span {
          width: 33.3%;
          color: #64b5f6;

          .removeTag {
            cursor: pointer;
            color: #666;
          }
        }
      }
      input {
        border: 1px solid #666;
        border-radius: 10px;
        background-color: #f5f5f5;
        font-size: 0.875rem;
        color: #333;
        text-align: center;
      }

      input:focus {
        outline: none;
      }
    }
  }
}
@media (max-width: 480px) {
  #new-posts {
    width: 100%;
    margin: auto;

    .contents {
      width: 98%;
      margin: auto;
    }
  }
}
</style>
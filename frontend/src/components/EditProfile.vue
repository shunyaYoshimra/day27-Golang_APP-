<template>
  <div id="edit-user">
    <div class="wrapper">
      <template v-if="!profileBool">
        <div class="profile">
          <p class="title">Profile</p>
          <textarea v-model="profile.description" placeholder="自己紹介をしよう(Write your profile)"></textarea>
          <div class="subject" v-for="(value, id) in subjects" :key="id">
            <label>
              <input v-model="profile.subject" :value="value" class="with-gap" name="group1" type="radio"  />
              <span>{{value}}</span>
            </label>
          </div>
          <label>*必須(Essential)</label><br>
          <button @click="createProfile()" class="btn blue lighten-3">go</button>
        </div>
         <hr>
      </template>
      <div class="contact">
        <template v-if="contactBool">
          <p class="title">Contact</p>
          <small class="explanation">Gaidai Chat Online は留学や就職で悩んだ際に、学生同士でコミュニケーションを取れるようにすることを目的としています。あなたの就活や留学の体験が気になった人からコンタクトを受けれるようにSNSとリンクさせましょう。</small><br><br>
          <label>
            <input v-model="contact.media" value="twitter" class="with-gap" name="group1" type="radio"  />
            <span>Twitter</span>
          </label>
          <label>
            <input v-model="contact.media" value="instagram" class="with-gap" name="group1" type="radio"  />
            <span>Instagram</span>
          </label>
          <br>
          <div class="account">
            <small>アカウント名を記入してください。Twitterの場合には最初の@を抜いて記入してください。(Please write your account name. If you choose Twiiter, please write it without @)
            </small>
            <input v-model="contact.account" type ="text">
          </div>
          <button @click="editContact()" class="btn blue lighten-3">go</button>
          <hr>
        </template>
      </div>
      <template v-if="abroadBool">
        <div class="abroad">
          <p class="title">Study Abroad</p>
          <small class="explanation">あなたの留学先や目的などを共有してみましょう。</small><br>
          <label for="country"><i class="material-icons tiny">place</i>Country</label><br>
          <input v-model="abroad.country" type="text" id="country"><br>
          <label for="college"><i class="material-icons tiny">school</i>College</label><br>
          <input v-model="abroad.college" type="text" id="college"><br>
          <label for="abroad-description"><i class="material-icons tiny">edit</i>Description</label><br>
          <textarea id="abroad-description" v-model="abroad.description"></textarea><br>
          <button @click="editAbroad()" class="btn blue lighten-3">go</button>
        </div>
      </template>
      <template v-else>
        <div class="abroad">
          <p class="title">Study Abroad</p>
          <small class="explanation">あなたの留学先や目的などを共有してみましょう。</small><br>
          <label for="country"><i class="material-icons tiny">place</i>Country</label><br>
          <input v-model="newAbroad.country" type="text" id="country"><br>
          <label for="college"><i class="material-icons tiny">school</i>College</label><br>
          <input v-model="newAbroad.college" type="text" id="college"><br>
          <label for="abroad-description"><i class="material-icons tiny">create</i>Description</label><br>
          <textarea id="abroad-description" v-model="newAbroad.description"></textarea><br>
          <button @click="createAbroad()" class="btn blue lighten-3">go</button>
        </div>
      </template>
      <hr>
      <template v-if="occupationBool">
        <div class="occupation">
          <p class="title">Occupation</p>
          <small class="explanation">あなたの仕事の様子や就活の様子を共有して、同じ業種を志望する後輩の参考になろう。</small><br>
          <div class="job-kinds">
            <label v-for="(value, id) in jobKinds" :key="id">
              <input v-model="occupation.kind" :value="value" class="with-gap" name="group2" type="radio"  />
              <span>{{value}}</span>
            </label>
          </div>
          <label for="company"><i class="material-icons tiny">work</i>Company(*任意)</label><br>
          <input v-model="occupation.company" type="text" id="company"><br>
          <label for="occupation-description"><i class="material-icons tiny">create</i>Description</label><br>
          <textarea v-model="occupation.description" id="occupation-description"></textarea><br>
          <button @click="editOccupation()" class="btn blue lighten-3">go</button>
        </div>
      </template>
      <template v-else>
        <div class="occupation">
          <p class="title">Occupation</p>
          <small class="explanation">あなたの仕事の様子や就活の様子を共有して、同じ業種を志望する後輩の参考になろう。</small><br>
          <div class="job-kinds">
            <label v-for="(value, id) in jobKinds" :key="id">
              <input v-model="newOccupation.kind" :value="value" class="with-gap" name="group2" type="radio"  />
              <span>{{value}}</span>
            </label>
          </div>
          <label for="company"><i class="material-icons tiny">work</i>Company(*任意)</label><br>
          <input v-model="newOccupation.company" type="text" id="company"><br>
          <label for="occupation-description"><i class="material-icons tiny">create</i>Description</label><br>
          <textarea v-model="newOccupation.description" id="occupation-description"></textarea><br>
          <button @click="createOccupation()" class="btn blue lighten-3">go</button>
        </div>
      </template>
    </div>
  </div>
</template>


<script>
import axios from "axios";
export default {
  data() {
    return {
      me: {},
      profileBool: true,
      contact: {},
      contactBool: true,
      subjects: ["英米(Engilish)", "イスパニア(Spanish)", "中国(Chinese)", "ロシア(Russian)", "国際関係(IR)", "2部英米(Engilish')", "院生(Graduate)", "留学生(Exchange Students)", "その他(Other)"],
      profile: {
        subject: "",
      },
      abroad: {},
      abroadBool: true,
      newAbroad: {
        country: "",
        college: "",
        description: "",
      },
      occupation: {},
      occupationBool: true,
      jobKinds: ["メーカー", "商社", "金融", "鉄道,航空", "運輸,物流", "不動産", "医療,福祉", "教育", "広告,出版,マスコミ", "小売", "公務員", "コンサルティング,調査", "ソフトウエア,インターネット"],
      newOccupation: {
        kind: "",
        company: "",
        description: ""
      }
    }
  },
  mounted() {
    // get me
    axios.get("/api/get_me").then((res) => {
      this.me = res.data.data;
    })
    // get user profile
    axios.get("/api/my_profile").then((res) => {
    }).catch((err) => {
      if (err.response.status === 404) {
        this.profileBool = false;
      }
    })
    // get user contact
    axios.get("/api/my_contact").then((res) => {
      this.contact = res.data.data;
    }).catch((err) => {
      if (err.response.status === 404) {
        this.contactBool = false;
      }
    });
    // get user abroad
    axios.get("/api/my_abroad").then((res) => {
      console.log(res.data)
      this.abroad = res.data.data;
    }).catch((err) => {
      if (err.response.status === 404) {
        this.abroadBool = false;
      }
    })
    // get user occupation
    axios.get("/api/my_occupation").then((res) => {
      this.occupation = res.data.data;
    }).catch((err) => {
      if (err.response.status === 404) {
        this.occupationBool = false;
      }
    })
  },
  methods: {
    createProfile() {
      const params = new URLSearchParams();
      params.append("description", this.profile.description)
      params.append("subject", this.profile.subject)
      axios.post("/api/v1/profiles/create", params).then((res) => {
        this.$router.push(`/profile/${this.me.id}`)
      }).catch((err) => {
        console.log(err.response);
      })
    },
    editContact() {
      const params = new URLSearchParams();
      params.append("media", this.contact.media);
      params.append("account", this.contact.account);
      axios.put("/api/contacts", params).then((res) => {
        this.$router.push(`/profile/${this.me.id}`)
      })
    },
    createAbroad() {
      const params = new URLSearchParams();
      params.append("country", this.newAbroad.country);
      params.append("college", this.newAbroad.college);
      params.append("description", this.newAbroad.description);
      axios.post("/api/abroads", params).then((res) => {
        this.$router.push(`/profile/${this.me.id}`)
      })
    },
    editAbroad() {
      const params = new URLSearchParams();
      params.append("country", this.abroad.country);
      params.append("college", this.abroad.college);
      params.append("description", this.abroad.description);
      axios.put("/api/abroads", params).then((res) => {
         this.$router.push(`/profile/${this.me.id}`)
      })
    },
    createOccupation() {
      const params = new URLSearchParams();
      params.append("kind", this.newOccupation.kind);
      params.append("company", this.newOccupation.company);
      params.append("description", this.newOccupation.description);
      axios.post("/api/occupations", params).then((res) => {
        this.$router.push(`/profile/${this.me.id}`)
      }).catch((err) => {
        console.log(err.response)
      })
    },
    editOccupation() {
      const params = new URLSearchParams();
      params.append("kind", this.occupation.kind);
      params.append("company", this.occupation.company);
      params.append("description", this.occupation.description);
      axios.put("/api/occupations", params).then((res) => {
        this.$router.push(`/profile/${this.me.id}`)
      }).catch((err) => {
        console.log(err.response)
      })
    }
  },
}
</script>

<style lang="scss" scoped>
#edit-user {
  background-color: #fff;
  margin-bottom: 40px;
  .wrapper {
    width: 600px;
    margin: auto;
    padding: 30px 0 30px;
    // common
    .title {
      font-weight: bold;
    }
    button {
      margin: 10px 0;
    }
    .explanation {
      font-weight: bold;
    }
    input {
      width: 400px;
      border: 1px solid #333;
    }
    input:focus {
      outline: none;
    }
    textarea {
      width: 400px;
      height: 180px;
    }
    textarea:focus {
      outline: none;
    }
    // profile
    .profile {
      textarea {
        width: 400px;
        height: 180px;
      }
      textarea:focus {
        outline: none;
      }
    }
    // contact
    .contact {
      margin-top: 50px;
      .account {
        margin-top: 15px;
      }
    }
    // abroad
    .abroad {
      margin-top: 50px;
    }
    // occupation
    .occupation {
      margin-top: 50px;
      .job-kinds {
        margin: 30px 0;
        display: flex;
        flex-wrap: wrap;
        label {
          width: 33.3%;
        }
      }
    }
  }
}
</style>
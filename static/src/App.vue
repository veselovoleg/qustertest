<template>
  <div id="app">
    <h2>The Questioner</h2>
    <ul>
      <table>
        <tr>
          <td colspan="3" id="nothing"></td>
          <td colspan="1" id="delete" title="Delete all answers" @click="deleteAnswers">Del</td>
        </tr>
        <tr v-for="(item, index) in data" v-bind:key="index">
          <td>{{item.question}}</td>
          <td class="answer" @click="choosedAnswer(index, 0)" title="Choose 'Yes'">{{answers[0]}}</td>
          <td class="answer" @click="choosedAnswer(index, 1)" title="Choose 'No'">{{answers[1]}}</td>
          <td v-if="choosed[index] == ''" class="nochecked"></td>
          <td v-if="choosed[index] == 'Yes'" class="checkedyes" title="You choosed 'Yes' in this row">Yes</td>
          <td v-if="choosed[index] ==  'No'" class="checkedno" title="You choosed 'No' in this row">No</td>
        </tr>
        <tr>
          <td colspan="4" id="sender-passive" v-if="!checkAnswers"></td>
          <td colspan="4" id="sender-active" v-if="checkAnswers" title="Send answers!" @click="sendAnswers">Send Answers?</td>
        </tr>
      </table>
    </ul>
  </div>
</template>

<script>
export default {
  name: "app",
  data() {
    return {
      data: [],
      answers: { 0: "Yes", 1: "No" },
      choosed: { 0: "", 1: "", 2: "" }
    };
  },
  mounted() {
    this.getData();
  },
  methods: {
    choosedAnswer(index, val) {
      this.$set(this.choosed, index, this.answers[val]);
    },
    deleteAnswers(){
      for (let i = 0; i < this.data.length; i++){
        this.$set(this.choosed, i, '');
      }
    },
    sendAnswers(){
      
    },
    getData() {
      axios
        .get("http://localhost:8081/getdata")
        .then(response => {
          this.data = response.data.items;
        })
        .catch(e => {
          console.log("Error GET data");
        });
    }
  },
  computed: {
    checkAnswers(){
      var ans = this.choosed;
      return ans[0] == '' || ans[1] == '' || ans[2] == '' ?  false : true;
    }
  }
};
</script>

<style>
body {
  background-color: #42b983;
}

#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #ddd;
  margin-top: 60px;
  width: 50%;
  margin: auto;
}

h1,
h2 {
  font-weight: normal;
  color: azure;
}

table {
  width: 90%;
  background-color: #2c3e50;
  border-collapse: collapse;
}

tr td {
  padding: 8px 16px;
}

.answer {
  background-color: #f7cc0ccc;
  width: 20%;
  color: aliceblue;
  cursor: pointer;
}

.answer:hover {
  background-color: #bea42f;
}

.nochecked,
.checkedyes,
.checkedno {
  width: 8%;
}

.checkedyes {
  background-color: green;
}

.checkedno {
  background-color: #ca1818;
}

.nochecked {
  background-color: rgb(5, 5, 5);
}

#sender-active, #sender-passive {
  height: 40px;
}

#sender-active {
  background-color: rgb(0, 26, 255);
  cursor: pointer;
}

#sender-active:hover {
  background-color: rgb(33, 44, 138);
}
#sender-passive {
  background-color: rgb(92, 92, 94);
}

#nothing{
  background-color: #42b983;
}

#delete{
  background-color: #ca1818;
}

#delete:hover{
  background-color: #b43232;
  cursor: pointer;
}

</style>

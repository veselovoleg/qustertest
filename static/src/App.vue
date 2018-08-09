<template>
  <div id="app">
    <h2>The Questioner</h2>
    <ul>
      <table v-for="(item, index) in data" v-bind:key="index">
        <tr>
          <td>{{item.question}}</td>
          <td class="answer" @click="choosedAnswer(index, 0)">{{answers[0]}}</td>
          <td class="answer" @click="choosedAnswer(index, 1)">{{answers[1]}}</td>
          <td v-if="choosed[index] == ''" class="nochecked">
            <!-- <span v-if="choosed[index] == 'Yes' && 'No'">Yes</span>
            <span v-if="choosed[index] == 'No'">No</span>
            <span v-if="choosed[index] == ''"></span> -->
          </td>
          <td v-if="choosed[index] == 'Yes' || choosed[index] ==  'No'" class="checked">

          </td>
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
      //let num = this.data[index].id;
      this.$set(this.choosed, index, this.answers[val]);
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
  /* computed: {
    checkAnswers() {
      var mass = [false, false, false];

      if(this.choosed[0] == false){
        mass[0] = false 
      } else {
        mass[0] = true 
      }
      this.choosed[0] == false ? mass[0] = false : mass[0] = true ;
      this.choosed[1] == false ? mass[1] = false : mass[1] = true ;
      this.choosed[2] == false ? mass[2] = false : mass[2] = true ;

      return mass;
    }
  } */
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

.checked, .nochecked {
  width: 5%
}

.checked{
  background-color: green
}

.nochecked{
  background-color: rgb(5, 5, 5)
}
</style>

<template>
  <div class="calculator">
    <div class="screen">{{ expression }}</div>
    <div class="key-board">
      <Button
        v-for="button in buttonList"
        :key="button.index"
        :text="button"
        @add-screen="addToScreen"
        @do-math="doMath"
        @clear-screen="clearScreen"
      />
    </div>
  </div>
</template>

<script>
import Button from "./components/Button.vue";
import axios from "axios";
export default {
  name: "App",
  components: {
    Button,
  },
  data() {
    return {
      expression: "",
    };
  },
  computed: {
    buttonList() {
      return [
        "C",
        "CE",
        "%",
        "/",
        "7",
        "8",
        "9",
        "*",
        "4",
        "5",
        "6",
        "-",
        "1",
        "2",
        "3",
        "+",
        "OCG",
        "0",
        ".",
        "=",
      ];
    },
  },
  methods: {
    DataToQuery(expresstion) {
      let operator = "";
      for (let i = 0; i < expresstion.length; i++) {
        let a = expresstion.charCodeAt(i); //convert to string adn check operator
        if (a < 48) {
          operator = expresstion[i];
          break;
        }
      }
      // operator
      //take value[0], from indexOf operator
      //from indexOf operator + 1, take value as length
      let indexOfOperator = expresstion.indexOf(operator);
      return (this.query =
        "/" +
        operator +
        "/" +
        expresstion.slice(0, indexOfOperator) +
        "/" +
        expresstion.slice(indexOfOperator + 1, expresstion.length));
    },
    async doMath() {
      console.log("sending");
      try {
        const response = await axios.get(
          "http://localhost:8000" + this.DataToQuery(this.expression)
        );
        this.expression = response.data.result;
      } catch (error) {
        console.log("can not get " + error);
      }
    },
    addToScreen(text) {
      return (this.expression += text);
    },
    clearScreen() {
      return (this.expression = ""), console.log("Cleared");
    },
  },
};
</script>

<style>
body {
  font-family: Roboto;
  background: #373b44;
  background: -webkit-linear-gradient(to right, #4286f4, #373b44);
  background: linear-gradient(to right, #4286f4, #373b44);
}
.calculator {
  width: 320px;
  height: 520px;
  background-color: #e965e2;
  top: 20px;
  margin: 0 auto;
  position: relative;
  border-radius: 5px;
  box-shadow: 0 4px 8px rgba(214, 105, 105, 0.2);
}

.screen {
  flex-direction: column;
  height: 120px;
  border-style: solid;
  border-radius: 4px;
  border-width: thin;
}

.screen {
  text-align: right;
  font-size: 50px;
}

.key_board {
  height: 400px;
}
.button {
  width: 50px;
  height: 50px;
  margin: 15px;
  float: left;
  border-radius: 10%;
  border-style: solid;
  font-weight: bold;
  font-size: 16px;
}

.button {
  cursor: pointer;
}

.button:hover {
  background-color: rgba(250, 58, 33, 0.952);
}
</style>

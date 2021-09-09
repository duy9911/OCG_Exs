<template>
  <button class="button" @click="handlerUserTyping">{{ text }}</button>
</template>
<script>
export default {
  name: "Button",
  props: { text: String },
  data() {
    return {
      theFirstInput: true,
    };
  },
  computed: {
    handlerUserTyping() {
      console.log(this.text);
      switch (this.text) {
        case "C":
        case "CE":
          return this.callClearScreen;
        case "=":
          return this.callDoMath;
        default:
          return this.callAddToScreen;
      }
    },
  },
  // actice event call respective func (active event goi den func tuong ung)
  methods: {
    callAddToScreen() {
      this.$emit("add-screen", this.text);
    },
    callDoMath() {
      this.$emit("do-math");
    },
    callClearScreen() {
      this.$emit("clear-screen");
    },
  },
};
</script>
<style scoped>
</style>
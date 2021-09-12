<template>
  <div class="container">
    <div id="banner">
      <img alt="Vue logo" src="./assets/imgs/illustration-hero.svg" />
      <div class="orderMenu">
        <h1>Order Summary</h1>
        <p class="quote">
          You can now listen to millions of songs, audiobooks, and podcasts on
          any device anywhere you like!
        </p>
        <div class="planning">
          <div class="icon-music">
            <img src="./assets/imgs/icon-music.svg" alt="" />
          </div>
          <div class="plan-list">
            <p class="title">{{ planName }} plan</p>
            <p class="price">{{ planPrice }}$/year</p>
          </div>
          <button @click="changePlan">Change</button>
        </div>
        <button class="status" @click="enterPlan">Proceed to Payment</button>
        <button class="cancel" @click="cancelPlan">Cancel</button>

        <div class="modal-container" v-if="this.status">
          <strong>Thank you!</strong> Enjoy!
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Order",
  props: {
    msg: String,
  },
  data() {
    return {
      planName: "",
      planPrice: null,
      indexOfplan: 0,
      status: false,
    };
  },
  //whenever the indexOfplan is changed watch will fun
  watch: {
    indexOfplan() {
      this.planName = this.planList[this.indexOfplan].planName;
      this.planPrice = this.planList[this.indexOfplan].planPrice;
    },
  },
  computed: {
    planList() {
      return [
        { planName: "Anual", planPrice: 2323 },
        { planName: "Premium", planPrice: 23223 },
        { planName: "Pro", planPrice: 2323 },
      ];
    },
  },

  created() {
    this.planName = this.planList[0].planName;
    this.planPrice = this.planList[0].planPrice;
  },

  methods: {
    changePlan() {
      this.indexOfplan++;
      if (this.indexOfplan >= this.planOptions.length) {
        this.indexOfplan = 0;
      }
    },
    cancelPlan() {
      this.status = false;
    },
    enterPlan() {
      this.status = true;
    },
  },
};
</script>

<style>
* {
  box-sizing: border-box;
}
button {
  cursor: pointer;
}
html {
  font-size: 14px;
  font-family: "Open Sans", sans-serif;
  color: #9babc0;
}
@media screen and (min-width: 768px) {
  body {
    background-image: url(./assets/imgs/pattern-background-desktop.svg);
    background-position: 0 0;
    background-repeat: no-repeat;
    background-size: 100% auto;
  }
}
@media screen and (max-width: 768px) {
  body {
    background-color: #d5e0fe;
  }
}
.container {
  width: 100%;
  height: 100%;
}
</style>

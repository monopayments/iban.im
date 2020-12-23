<template>
  <div>
    <div id="scan">
      <button v-on:click="recognize">recognize</button>
      <hr />
      <input type="file" @change="onFileChange" />

      <div id="preview">
        <img id="iban_img" v-if="url" :src="url" />
      </div>
      <hr />
      <div>
        bulunan metin:
        <div id="iban_raw"></div>
      </div>
      <div>
        bulunan iban:
        <strong id="iban"></strong>
      </div>
    </div>
  </div>
</template>

<script>
/* eslint-disable */
import { createWorker, PSM, OEM } from "tesseract.js";
const worker = createWorker({
  logger: (m) => console.log(m),
});

export default {
  name: "Scan",
  data() {
    return {
      url: null,
      text: null,
      iban: null,
    };
  },
  methods: {
    recognize: async () => {
      const iban_img = document.getElementById("iban_img");
      const iban_raw = document.getElementById("iban_raw");
      const iban_el = document.getElementById("iban");
      await worker.load();
      await worker.loadLanguage("eng");
      await worker.initialize("eng", OEM.LSTM_ONLY);
      await worker.setParameters({
        tessedit_pageseg_mode: PSM.SINGLE_BLOCK,
      });
      const {
        data: { text },
      } = await worker.recognize(iban_img);
      console.log(text)

      let str_iban = text;
      iban_raw.innerHTML = str_iban;
      str_iban = str_iban.match(
        /[a-zA-Z]{2}[0-9]{2}[a-zA-Z0-9]{4}[0-9]{7}([a-zA-Z0-9]?){0,16}/m
      );

      iban_el.innerHTML = str_iban;
    },
    onFileChange(e) {
      const file = e.target.files[0];
      this.url = URL.createObjectURL(file);
    },
  },
};
</script>
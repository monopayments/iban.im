<template>
  <div class="box">
    <div id="scan">
      <ul class="iconList">
        
        <li style="--color: var(--primary-2)">
          <div id="preview">
            <img id="iban_img" style="max-height: 40px;" alt="" v-if="url" :src="url" />
          </div>
        </li>
        <li>
          <input type="file" @change="onFileChange" />
        </li>
        <template v-if="scanning">
          <v-progress-circular
              indeterminate
              color="primary"
          />
        </template>
        <template v-else>
          <li>
            <button v-on:click="recognize" class="btn">recognize</button>
          </li>
          <li style="--color: var(--primary-3)">
            Text in the image:
            <div id="iban_raw">{{ibanRaw}}</div>
          </li>
          <li style="--color: var(--primary-3)">
            possible IBAN:
            <strong id="iban">{{iban}}</strong>
          </li>
        </template>

      </ul>
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
  data () {
    return {
      url: null,
      text: null,
      iban: null,
      ibanRaw : null,
      scanning : false,
    };
  },
  methods: {
    async recognize() {
      this.scanning = true;
      const iban_img = document.getElementById("iban_img");
      await worker.load();
      await worker.loadLanguage("eng");
      await worker.initialize("eng", OEM.LSTM_ONLY);
      await worker.setParameters({
        tessedit_pageseg_mode: PSM.SINGLE_BLOCK,
      });
      const {
        data: { text },
      } = await worker.recognize(iban_img);
      this.scanning = false;
      console.log(text);
      this.ibanRaw = text;
      this.iban = text.match(
        /[a-zA-Z]{2}[0-9]{2}[a-zA-Z0-9]{4}[0-9]{7}([a-zA-Z0-9]?){0,16}/m
      );
    },
    onFileChange(e) {
      const file = e.target.files[0];
      this.url = URL.createObjectURL(file);
    },
  },
};
</script>
<template>
  <main>
    <div id="result" class="result">{{ data.resultText }}</div>
    <div class="prolabore">
      <div class="input-group">
        <BaseInput v-model="data.pretendedSalary" label="Pretenção salarial" type="number"/>
        <BaseInput v-model="data.workDayPerWeek" label="Dias de trabalho por semana" type="number"/>
        <BaseInput v-model="data.hoursPerDay" label="Horas trabalhadas por dia" type="number"/>
      </div>
      <button class="btn" @click="prolabore">Calcular</button>
    </div>
  </main>
</template>

<script setup>
  import { reactive } from "vue"
  import { Prolabore } from "../../wailsjs/go/main/App"
  import BaseInput from "./BaseInput.vue"

  const data = reactive({
    workDayPerWeek: null,
    hoursPerDay: null,
    pretendedSalary: null,
    resultText: "Por favor entre as informações abaixo 👇",
  })

  async function prolabore() {
    const workDayPerWeek = parseInt(data.workDayPerWeek)
    const hoursPerDay = parseInt(data.hoursPerDay)
    const pretendedSalary = parseInt(data.pretendedSalary)
    data.resultText = await Prolabore(workDayPerWeek, hoursPerDay, pretendedSalary)
  }
</script>

<style scoped>
  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .prolabore {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-between;

    .input-group {
      display: flex;
      gap: 1rem;
    }

    .btn {
      border-radius: 0.2rem;
      border: none;
      margin: 1rem 0 0 0;
      padding: 0.5rem 1rem;
      cursor: pointer;
    }

    .btn:hover {
      background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
      color: #333333;
    }
  }
</style>

<script setup>
import {reactive} from 'vue'
import {Prolabore} from '../../wailsjs/go/main/App'

const data = reactive({
  workDayPerWeek: null,
  hoursPerDay: null,
  pretendedSalary: null,
  resultText: "Por favor entre as informações abaixo 👇",
})

function prolabore() {
  Prolabore(data.workDayPerWeek, data.hoursPerDay, data.pretendedSalary).then(result => {
    data.resultText = result
  })
}

</script>

<template>
  <main>
    <div id="result" class="result">{{ data.resultText }}</div>
    <div class="prolabore">
      <div id="inputGroup" class="input-group">
        <div id="inputWorkDayPerWeek" class="input-box">
          <label for="workDayPerWeek">Dias de trabalho por semana</label>
          <input id="workDayPerWeek" v-model="data.workDayPerWeek" autocomplete="off" class="input" type="number" placeholder="0"/>
        </div>
        <div id="inputHoursPerDay" class="input-box">
          <label for="workDayPerWeek">Horas trabalhadas por dia</label>
          <input id="hoursPerDay" v-model="data.hoursPerDay" autocomplete="off" class="input" type="number" placeholder="0"/>
        </div>
        <div id="inputPretendedSalary" class="input-box">
          <label for="workDayPerWeek">Pretenção salarial</label>
          <input id="pretendedSalary" v-model="data.pretendedSalary" autocomplete="off" class="input" type="number" placeholder="0"/>
        </div>
      </div>
      <button class="btn" @click="prolabore">Calcular</button>
    </div>
  </main>
</template>

<style scoped>
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}

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
}

.prolabore .input-group {
  display: flex;
  justify-content: space-between;
  gap: 1rem;
}

.prolabore .input-box {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  gap: 1rem;
}

.prolabore .btn {
  border-radius: 0.2rem;
  border: none;
  margin: 1rem 0 0 0;
  padding: 0.5rem 1rem;
  cursor: pointer;
}

.prolabore .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.prolabore .input-box .input {
  border: none;
  border-radius: 0.2rem;
  outline: none;
  height: 2rem;
  padding: 0 0.5rem;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.prolabore .input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.prolabore .input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>

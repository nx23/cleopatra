<template>
  <main>
    <v-form v-model="valid">
      <v-container fluid>
        <v-row>
          <v-col>
            <div class="result">{{ resultText }}</div>
          </v-col>
        </v-row>
        <v-row>
          <v-col
            cols="12"
            md="4"
          >
            <v-text-field
              class="input-box"
              v-model="salary"
              :rules="requiredField"
              label="Salário esperado"
              variant="outlined"
            />
          </v-col>

          <v-col
            cols="12"
            md="4"
          >
            <v-text-field
              v-model="hoursPerDay"
              :rules="workHoursRules"
              label="Horas trabalhadas por dia"
              variant="outlined"
            />
          </v-col>

          <v-col
            cols="12"
            md="4"
          >
            <v-text-field
              v-model="workDaysPerWeek"
              :rules="weekdaysRules"
              label="Dias de trabalho por semana"
              variant="outlined"
            />
          </v-col>
        </v-row>
        <v-row>
          <v-col
            cols="12"
            md="4"
          ></v-col>
        </v-row>
        <v-row class="debug">
          <div class="charge-header">Adicionar Gastos</div>
        </v-row>
        <v-row>
          <v-col
            cols="12"
            md="4"
          >
            <v-text-field
              v-model="chargeName"
              :rules="requiredField"
              label="Nome do Gasto"
              variant="outlined"
            />
          </v-col>
          <v-col
            cols="12"
            md="4"
          >
            <v-text-field
              v-model="chargeCost"
              :rules="requiredField"
              label="Valor do Gasto"
              variant="outlined"
            />
          </v-col>
        </v-row>
        <v-row>
          <v-data-table :items="items" :headers="headers" class="transparent-table"></v-data-table>
        </v-row>
        <v-row>
          <v-col>
            <v-btn class="calcular" prepend-icon="mdi-cash-check" @click="prolabore">
              Calcular
            </v-btn>
          </v-col>
        </v-row>
      </v-container>
    </v-form>
  </main>
</template>

<script>
  import { Prolabore } from "../../wailsjs/go/main/App"

  export default {
    data: () => ({
      valid: false,
      salary: null,
      hoursPerDay: null,
      workDaysPerWeek: null,
      chargeName: null,
      chargeCost: null,
      resultText: "Por favor entre as informações abaixo 👇",
      requiredField: [
        value => {
          if (value) return true

          return 'O campo é obrigatório.'
        },
      ],
      weekdaysRules: [
        value => {
          if (value) return true

          return 'O campo é obrigatório.'
        },
        value => {
          if (value <= 7) return true

          return 'O limite é 7 dias.'
        },
      ],
      workHoursRules: [
        value => {
          if (value) return true

          return 'O campo é obrigatório.'
        },
        value => {
          if (value <= 24) return true

          return 'O limite é 24 horas.'
        },
      ],
      prolabore: async function() {
        const workDaysPerWeek = parseInt(this.workDaysPerWeek)
        const hoursPerDay = parseInt(this.hoursPerDay)
        const pretendedSalary = parseInt(this.salary)
        this.resultText = await Prolabore(workDaysPerWeek, hoursPerDay, pretendedSalary)
      },
      headers:[
        { title: "Tipos de Gastos", value: "type" },
        { title: "Valor", value: "value" },
      ],
      items: [],
    }),
    created() {
      this.prolabore = this.prolabore.bind(this);
    },
  }
</script>

<style scoped>
  .v-form {
    background-color: rgba(255, 255, 255, 0.8);
    color: black;
  }

  .charge-header {
    color: black;
    display: flex;
    justify-content: flex-start;
  }
</style>
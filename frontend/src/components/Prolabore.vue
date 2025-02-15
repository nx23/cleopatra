<template>
  <main>
    <v-container fluid>
      <v-form v-model="valid" class="income-form">
        <v-row>
          <v-col>
            <div>Por favor entre as informações abaixo 👇</div>
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" md="4">
            <v-text-field
              class="input-box"
              v-model="salary"
              :rules="requiredField"
              label="Salário esperado"
              variant="outlined"
            />
          </v-col>

          <v-col cols="12" md="4">
            <v-text-field
              v-model="hoursPerDay"
              :rules="workHoursRules"
              label="Horas trabalhadas por dia"
              variant="outlined"
            />
          </v-col>

          <v-col cols="12" md="4">
            <v-text-field
              v-model="workDaysPerWeek"
              :rules="weekdaysRules"
              label="Dias de trabalho por semana"
              variant="outlined"
            />
          </v-col>
        </v-row>
        <v-row>
          <v-col cols="12" md="4"></v-col>
        </v-row>
      </v-form>
      <v-form v-model="validCharge" class="charge-form">
        <v-row>
          <div class="charge-header">Gastos</div>
        </v-row>
        <v-row>
          <v-col md="4">
            <v-text-field
              v-model="chargeName"
              :rules="requiredField"
              label="Nome do Gasto"
              variant="outlined"
            />
          </v-col>
          <v-col md="4">
            <v-text-field
              v-model="chargeCost"
              :rules="requiredField"
              label="Valor do Gasto"
              variant="outlined"
            />
          </v-col>
        </v-row>
        <v-row>
          <v-col class="add-charge">
            <v-btn prepend-icon="mdi-plus" @click="addCharge">Adicionar Gasto</v-btn>
          </v-col>
        </v-row>
        <v-row>
          <v-data-table :items="items" :headers="headers" class="transparent-table">
            <template v-slot:item.actions="{ item }">
              <v-btn class="action-button"
                v-for="(action, index) in item.actions"
                :key="index"
                @click="action.onClick(item)"
                variant="plain"
              >
                <v-icon>{{ action.icon }}</v-icon>
              </v-btn>
            </template>
          </v-data-table>
        </v-row>
      </v-form>
      <v-row>
        <v-col>
          <v-btn class="calcular" prepend-icon="mdi-cash-check" @click="prolabore">
            Calcular
          </v-btn>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <div>{{ resultText }}</div>
        </v-col>
      </v-row>
    </v-container>
  </main>
</template>

<script>
  import { Prolabore } from "../../wailsjs/go/main/App"

  export default {
    data: () => ({
      valid: false,
      validCharge: false,
      salary: null,
      hoursPerDay: null,
      workDaysPerWeek: null,
      chargeName: null,
      chargeCost: null,
      resultText: "",
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
        const charges = this.items.reduce((total, item) => total + parseInt(item.value), 0)
        this.resultText = await Prolabore(workDaysPerWeek, hoursPerDay, pretendedSalary, charges)
      },
      addCharge: function() {
        if (!this.validCharge) return
        
        this.validCharge = false
        this.items.push({
          type: this.chargeName,
          value: this.chargeCost,
          actions: [
            {
              icon: "mdi-delete",
              onClick: this.removeCharge,
            },
          ],
        })

        this.chargeName = null
        this.chargeCost = null
      },
      removeCharge: function(chargeName) {
        this.items = this.items = this.items.filter(item => item.type !== chargeName.type)
      },
      items: [],
      headers: [
        { title: "Tipos de Gastos", key: "type" },
        { title: "Valor", key: "value", align: "center" },
        { title: "Ações", key: "actions", align: "center", sortable: false },
      ],
    }),
    created() {
      this.prolabore = this.prolabore.bind(this);
      this.addCharge = this.addCharge.bind(this);
      this.removeCharge = this.removeCharge.bind(this);
    },
  }
</script>

<style scoped>
  .income-form {
    background-color: rgba(255, 255, 255, 0.5);
    padding: 0 1rem 0 1rem;
    margin: 0 -1rem 0 -1rem;
  }

  .charge-form {
    padding: 1rem 1rem 0 1rem;
    background-color: rgba(255, 255, 255, 0.5);
    margin: 2rem -1rem 0 -1rem;
  }

  .v-container {
    background-color: rgba(255, 255, 255, 0.5);
    color: black;
    padding-top: 1rem;
    margin-top: -0.2rem;
    height: 100vh;
  }

  .v-container {
    .v-row {
      margin-top: -1rem;
    }

    .v-data-table {
      margin-top: 1rem;
    }

    .v-btn {
      margin-top: 1rem;
    }
  }

  .charge-header {
    color: black;
    display: flex;
    justify-content: flex-start;
    padding: 1rem;
  }

  .transparent-table {
    display: flex;
    background-color: rgba(255, 255, 255, 0.7);

    .action-button {
      display: flex;
      justify-self: center;
      margin-top: -0.5rem;
    }
  }

  .add-charge {
    display: flex;
    justify-content: center;
    margin-top: -1rem;
  }
</style>
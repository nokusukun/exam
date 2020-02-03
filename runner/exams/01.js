/**
 * prop:Name
 * prop:Activity
 */

function campaignInput(days, bakers, cakes, waffles, pancakes) {
  let foodMoney = cakes * 45 + waffles * 5.8 + pancakes * 3.2;
  let grossFoodMoney = foodMoney * bakers * days;
  let costNumber = grossFoodMoney / 8;
  let totalRaised = grossFoodMoney - costNumber;
  console.log(totalRaised.toFixed(2));
}

const input1 = process.argv[2]
const input2 = process.argv[3]
const input3 = process.argv[4]
const input4 = process.argv[5]
const input5 = process.argv[6]

campaignInput(input1, input2, input3, input4, input5);
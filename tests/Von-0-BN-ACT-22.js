const input1 = process.argv[2]
const input2 = process.argv[3]
const input3 = process.argv[4]
const input4 = process.argv[5]
const input5 = process.argv[6]

function charityCampaign(days, bakers, cakes, waffles, pancakes) {
    let profitCakes=cakes*45;
    let profitWaffles=waffles*5.80;
    let profitPancakes=pancakes*3.20;
    let amountPerDay=(profitCakes+profitWaffles+profitPancakes)*bakers;
    let amountForCampaign=amountPerDay*days;
    let amountAfterCosts=amountForCampaign*.875;
    console.log(amountAfterCosts.toFixed(2));
}

charityCampaign(input1, input2, input3, input4, input5);
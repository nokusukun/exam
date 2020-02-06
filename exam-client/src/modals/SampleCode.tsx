import React from "react";
import DialogTitle from "@material-ui/core/DialogTitle";
import Dialog from "@material-ui/core/Dialog";
import { DialogContent } from "@material-ui/core";

export interface SimpleDialogProps {
  open: boolean;
}

const testCode = `
// Name: Von, Villamor E.
// Activity: 0-BN-ACT-22

const input1 = process.argv[2];
const input2 = process.argv[3];
const input3 = process.argv[4];
const input4 = process.argv[5];
const input5 = process.argv[6];

function charityCampaign(days, bakers, cakes, waffles, pancakes) {
  let profitCakes = cakes * 45;
  let profitWaffles = waffles * 5.8;
  let profitPancakes = pancakes * 3.2;
  let amountPerDay = (profitCakes + profitWaffles + profitPancakes) * bakers;
  let amountForCampaign = amountPerDay * days;
  let amountAfterCosts = amountForCampaign * 0.875;
  console.log(amountAfterCosts.toFixed(2));
}

charityCampaign(input1, input2, input3, input4, input5);
`;

export default function SampleCode(props: SimpleDialogProps) {
  const { open } = props;

  return (
    <Dialog aria-labelledby="sample-code" open={open} fullWidth>
      <DialogTitle id="sample-code">Sample Code</DialogTitle>
      <DialogContent>
        <h3>CHARITY-CAMPAIGN-0-BN</h3>
        <pre>{testCode}</pre>
      </DialogContent>
    </Dialog>
  );
}

import React, { useRef, useState, useEffect } from "react";
import Axios from "axios";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import TextField from "@material-ui/core/TextField";
import CheckCircleOutlineIcon from "@material-ui/icons/CheckCircleOutline";
import ErrorOutlineIcon from "@material-ui/icons/ErrorOutline";
import red from "@material-ui/core/colors/red";
import green from "@material-ui/core/colors/green";

import "./App.css";
import {
  Button,
  FormControl,
  InputLabel,
  Select,
  MenuItem
} from "@material-ui/core";
import SampleCode from "./modals/SampleCode";

interface Activity {
  id: string;
  description: string;
}

interface Output {
  passed: boolean;
  inputs: string[];
  expectedOutput: string;
  actualOutput: string;
}

const App = () => {
  const labelRef = React.createRef<HTMLLabelElement>();
  const [labelWidth, setLabelWidth] = useState(0);
  const [activities, setActivities] = useState([
    { id: "", description: "" }
  ] as Activity[]);
  const [activity, setActivity] = useState("");
  const [code, setCode] = useState("");
  const [output, setOutput] = useState([] as Output[]);

  const apiRoute = process.env.NODE_ENV === 'development' ? 'http://localhost:8888' : ''

  useEffect(() => {
    if (labelRef.current) {
      setLabelWidth(labelRef.current.offsetWidth);
    }

    (async () => {
      const activityRequest = await Axios.get(
        `${apiRoute}/activities`
      );
      console.log(activityRequest.data);
      setActivities(activityRequest.data);
    })();
  }, []);

  const submitCode = async () => {
    const body = JSON.stringify({
      submitterID: new Date().getTime().toString(),
      activityID: activity,
      code
    });

    const codeSubmitRequest = await Axios.post(
      `${apiRoute}/exam/submit`,
      body
    );
    setOutput(codeSubmitRequest.data as Output[]);
  };

  const generateReport = (outputs: Output[]) => {
    return outputs.map((out, index) => {
      return (
        <div style={{ marginBottom: "10px", textAlign: "left" }}>
          <Grid container spacing={1}>
            <Grid item xs={12}>
              {out.passed ? (
                <span style={{ color: "green" }}>
                  <CheckCircleOutlineIcon /> Test Passed
                </span>
              ) : (
                <span style={{ color: "red" }}>
                  <ErrorOutlineIcon /> Test Failed
                </span>
              )}{" "}
              ({`${index + 1} of ${outputs.length}`})
            </Grid>
            <Grid item xs={6}>
              Actual
            </Grid>
            <Grid item xs={6}>
              {out.actualOutput}
            </Grid>
            <Grid item xs={6}>
              Expected
            </Grid>
            <Grid item xs={6}>
              {out.expectedOutput}
            </Grid>
          </Grid>
        </div>
      );
    });
  };

  const [open, setOpen] = React.useState(false);

  const handleClickOpen = () => {
    setOpen(true);
  };

  return (
    <div className="App">
      <SampleCode open={open} />

      <Container style={{ marginTop: "10px" }}>
        <Grid container spacing={3}>
          <Grid item xs={6}>
            <Paper style={{ padding: "10px" }}>
              <FormControl
                variant="outlined"
                fullWidth
                style={{ marginBottom: "10px", textAlign: "left" }}
              >
                <InputLabel ref={labelRef} id="exam-type-label">
                  Activity ID
                </InputLabel>
                <Select
                  labelId="exam-type-label"
                  id="exam-type"
                  value={activity}
                  onChange={event => {
                    setActivity(event.target.value as string);
                  }}
                  labelWidth={labelWidth}
                >
                  {activities.map((act: Activity) => {
                    return <MenuItem value={act.id}>{act.id}</MenuItem>;
                  })}
                </Select>
              </FormControl>
              <TextField
                id="submission-code"
                label="Code"
                multiline
                rows="4"
                rowsMax="20"
                placeholder="Code"
                variant="outlined"
                fullWidth
                onChange={evt => {
                  setCode(evt.target.value);
                }}
              />
              <Button
                variant="outlined"
                color="default"
                style={{ marginTop: "10px", marginRight: "5px" }}
                onClick={() => handleClickOpen()}
              >
                Sample
              </Button>
              <Button
                variant="outlined"
                color="primary"
                style={{ marginTop: "10px" }}
                onClick={() => submitCode()}
              >
                Check
              </Button>
            </Paper>
          </Grid>
          <Grid item xs={6} alignContent="flex-start">
            <div style={{ padding: "10px" }}>
              <h1>Results</h1>
              <hr />
              {output.length <= 0
                ? "Waiting for submission..."
                : generateReport(output)}
              {/* <TextField
                id="Results"
                label="Results"
                multiline
                rows="4"
                rowsMax="20"
                placeholder="Results"
                variant="outlined"
                fullWidth
                value={output}
              /> */}
            </div>
          </Grid>
        </Grid>
      </Container>
    </div>
  );
};

export default App;

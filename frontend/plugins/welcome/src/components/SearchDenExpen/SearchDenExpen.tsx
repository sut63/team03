import React, { useState } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import ExitToAppRoundedIcon from '@material-ui/icons/ExitToAppRounded';
import moment from 'moment';
import { Page, pageTheme, Header, Content} from '@backstage/core';
import { Grid, Button, TextField, Typography, FormControl } from '@material-ui/core';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
import { EntDentalexpense } from '../../api';
import { Alert } from '@material-ui/lab';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    headsearch: {
      width: 'auto',
      margin: '10px',
      color: '#FFFFFF',
      background: '#2196F3',
    },
    margin: {
      margin: theme.spacing(1),
    },
    margins: {
      margin: theme.spacing(2),
    },

    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
    paper: {
      marginTop: theme.spacing(3),
      marginBottom: theme.spacing(3),
    },
    table: {
      minWidth: 500,
    },


  }),
);

export default function ComponentsTable() {
  const classes = useStyles();

  const [tax, setTax] = useState(String);
  const [dentalexpense, setDentalexpense] = useState<EntDentalexpense[]>([])
  const [alert, setAlert] = useState(true);
  const [status, setStatus] = useState(false);




  const Taxhandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setStatus(false);
    setTax(event.target.value as string);
  };

  const cleardata = () => {
    setTax("");
    setStatus(false);
    setDentalexpense([]);

  }

  

  const SearchDentalexpense = async () => {
    setStatus(true);
    setAlert(true);
    const apiUrl = `http://localhost:8080/api/v1/searchdentalexpenses?dentalexpense=${tax}`;
    const requestOptions = {
        method: 'GET',
    };
    fetch(apiUrl, requestOptions)
        .then(response => response.json())
        .then(data => {
            console.log(data.data)
            setStatus(true);
            setAlert(false);
            setDentalexpense([]);

          if (data.data != null) {
                if(data.data.length >= 1) {
                  setStatus(true);
                  setAlert(true);
                  console.log(data.data)
                  setDentalexpense(data.data);
                }
            }
        });

}

  return (

    <Page theme={pageTheme.service}>
      <Header
        title={`Dental System`}
        subtitle="ค้นหารายการค่ารักษา">
        <table>
        <Button variant="contained" color="default" href="/" startIcon={<ExitToAppRoundedIcon />}> Logout
        </Button>
        </table>

      </Header>
      <Content>
        <Grid container item xs={12} justify="center">
          <Grid item xs={5}>
            <Paper>

              <Typography align="center" >
                <div style={{ background: "#709fb0", height: 50 }}>
                  <h1 style={
                    {
                      color: "#FFFFFF",
                      borderRadius: 5,
                      height: 18,
                      padding: '0 30px',
                      fontSize: '30px',
                    }}>
                    ค้นหารายการค่ารักษา
            </h1>
                </div>

                <div>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >
                    <div className={classes.paper}>เลขที่กำกับภาษี</div>
                    <TextField
                      id="tax"
                      value={tax}
                      onChange={Taxhandlehange || ''}
                      type="string"
                      size="medium"

                      style={{ width: 300 }}
                    />
                  </FormControl>
                </div>

                <div></div>
                
                <Button
                  onClick={() => {
                    SearchDentalexpense();
                  }}

                  endIcon={<SearchTwoToneIcon />}
                  className={classes.margins}
                  variant="contained"
                  style={{ background: "#A1C4D8", height: 40 }}>
                  <h3
                    style={
                      {
                        color: "#FFFFFF",
                        padding: '0 10px',

                      }
                    }>
                    ค้นหาข้อมูล
            </h3>
                </Button>
                <Button
                  onClick={() => {
                    cleardata();

                  }}
                  className={classes.margins}
                  variant="contained"
                  style={{ background: "#3792cb", height: 40 }}>
                  <h3
                    style={
                      {
                        color: "#c6a9a3",
                        padding: '0 25px',

                      }
                    }>
                    เคลียร์ข้อมูล
            </h3>
                </Button>
              </Typography>
            </Paper>

            {status ? (
              <div>
                {alert ? (
                  <Alert severity="success">
                    แสดงรายการค่ารักษา
                  </Alert>
                )
                  : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    ไม่พบข้อมูล
                  </Alert>
                  )}
              </div>
            ) : null}

          </Grid>
        </Grid>
        

        
        <Grid  className={classes.paper}>
        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="left">ลำดับที่</TableCell>
                <TableCell align="left">เลขที่กำกับภาษี</TableCell>
                <TableCell align="left">ชื่อผู้ชำระค่าบริการ</TableCell>
                <TableCell align="left">บริการทันตกรรม</TableCell>
                <TableCell align="left">ค่าบริการ</TableCell>
                <TableCell align="left">ประเภทการชำระ</TableCell>
                <TableCell align="left">วันที่และเวลาในการชำระ</TableCell>
                </TableRow>
            </TableHead>
          <TableBody>
            {dentalexpense.map((item: any) => (
                <TableRow key={item.id}>
                  <TableCell align="left">{item.id}</TableCell>
                  <TableCell align="left">{item.Tax}</TableCell>
                  <TableCell align="left">{item.Name}</TableCell>
                  <TableCell align="left">{item.edges?.Medicalfile?.Detial}</TableCell>
                  <TableCell align="left">{item.Amount}</TableCell>
                  <TableCell align="left">{item.edges?.Pricetype?.name}</TableCell>
                  <TableCell align="left">{moment(item.AddedTime).format('DD/MM/YYYY HH.mm น.')}</TableCell>
                                    
                </TableRow>
                ))}
            </TableBody>
            </Table>
        </TableContainer>
        
        </Grid>

      </Content>
    </Page>
  );

}
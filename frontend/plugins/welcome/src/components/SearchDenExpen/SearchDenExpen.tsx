import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';

import { EntDentalexpense } from '../../api/models/EntDentalExpense';

import Swal from 'sweetalert2'
import { Link as RouterLink } from 'react-router-dom';
import moment from 'moment';
import { Page, pageTheme, Header, Content, Link } from '@backstage/core';
import { Grid, Button, TextField, Typography, FormControl } from '@material-ui/core';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';


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
const Toast = Swal.mixin({
  // toast: true,
  position: 'center',
  showConfirmButton: false,
  //timer: 3000,
  //timerProgressBar: true,
  showCloseButton: true,

});


export default function ComponentsTable() {

  
  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState(false);

  //---------------------------
  const [checktax, setcheckTaxs] = useState(false);
  const [dentalexpense, setDentalexpense] = useState<EntDentalexpense[]>([])

  //--------------------------
  const [tax, setTaxs] = useState(String);
  const profile = { givenName: 'ระบบค้นหารายการค่ารักษา' };
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  useEffect(() => {
    const getDentalexpense = async () => {
      const res = await api.listDentalexpense({ offset: 0 });
      setLoading(false);
      setDentalexpense(res);
    };
    getDentalexpense();
  }, [loading]);

  //-------------------
  const Taxhandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(false);
    setcheckTaxs(false);
    setTaxs(event.target.value as string);

  };

  const cleardata = () => {
    setTaxs("");
    setSearch(false);
    setcheckTaxs(false);
    setSearch(false);

  }
  //---------------------
  const checkdentalexpense = async () => {
    var check = false;
    dentalexpense.map(item => {
      if (tax != "") {
        if (item.tax?.includes(tax)) {
          setTaxs(true);
          alertMessage("success", "พบข้อมูลที่ค้นหา");
          check = true;
        }
      }
    })
    if (!check) {
      alertMessage("error", "ไม่พบข้อมูลที่ค้นหา");
    }
    console.log(checktax)
    if (tax == "") {
      alertMessage("info", "แสดงข้อมูลงานวิจัยทั้งหมดในระบบ");
    }
  };

  return (

    <Page theme={pageTheme.home}>
      <Header
        title={`${profile.givenName}`}
      >
        <table>
          <tr>
           
            <th>
              &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
          <Link component={RouterLink} to="/Menu">
                <Button variant="contained" style={{ background: 'linear-gradient(45deg, #3399FF 15%, #9900FF 120%)', height: 36 }}>
                  <h3
                    style={
                      {
                        color: "#FFFAF0",
                        borderRadius: 10,
                        height: 25,
                        padding: '0 20px',
                      }
                    }>
                    กลับหน้าหลัก
            </h3>
                </Button>
              </Link>
            </th>
          </tr>
        </table>

      </Header>
      <Content>
        <Grid container item xs={12} justify="center">
          <Grid item xs={5}>
            <Paper>

              <Typography align="center" >
                <div style={{ background: 'linear-gradient(45deg, #00FFCC 15%, #9900FF 120%)', height: 55 }}>
                  <h1 style={
                    {
                      color: "#000000",
                      borderRadius: 5,
                      height: 18,
                      padding: '0 30px',
                      fontSize: '40px',
                    }}>
                    ค้นหารายการค่ารักษา
            </h1>
                </div>

                <div>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >
                    <div className={classes.paper}><strong>เลขกำกับภาษี</strong></div>
                    <TextField
                      id="tax"
                      value={tax}
                      onChange={Taxhandlehange}
                      type="string"
                      size="small"

                      style={{ width: 200 }}
                    />
                  </FormControl>
                </div>
                <div></div>
                <Button
                  onClick={() => {
                    checkdentalexpense();
                    setSearch(true);

                  }}
                  endIcon={<SearchTwoToneIcon />}
                  className={classes.margins}
                  variant="contained"
                  style={{ background: "#9900FF", height: 40 }}>
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
                  style={{ background: "#9900FF", height: 40 }}>
                  <h3
                    style={
                      {
                        color: "#FFFFFF",
                        padding: '0 25px',

                      }
                    }>
                    เคลียร์ข้อมูล
            </h3>
                </Button>
              </Typography>
            </Paper>
          </Grid>
        </Grid>


        <Grid container justify="center">
          <Grid item xs={12} md={10}>
            <Paper>
              {search ? (
                <div>
                  {  checktax ? (
                    <TableContainer component={Paper}>
                      <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                          <TableRow>
                            <TableCell align="center">ลำดับที่</TableCell>
                            <TableCell align="center">ชื่องานวิจัย</TableCell>
                            <TableCell align="center">จำนวนหน้า</TableCell>
                            <TableCell align="center">ปีที่พิมพ์</TableCell>
                            <TableCell align="center">ชื่อผู้แต่ง</TableCell>
                            <TableCell align="center">ประเภทงานวิจัย</TableCell>
                            <TableCell align="center">วันที่</TableCell>
                          </TableRow>
                        </TableHead>
                        <TableBody>

                          {dentalexpense.filter((filter: any) => filter.tax.includes(tax)).map((item: any) => (
                            <TableRow key={item.id}>
                              <TableCell align="center">{item.id}</TableCell>
                              <TableCell align="center">{item.tax}</TableCell>
                              <TableCell align="center">{item.rates}</TableCell>
                              <TableCell align="center">{item.name}</TableCell>
                              <TableCell align="center">{item.edges?.medicalfile?.edges?.Patient?.name}</TableCell>
                              <TableCell align="center">{item.edges?.pricetypes?.name}</TableCell>
                              <TableCell align="center">{moment(item.AddedTime).format('DD/MM/YYYY HH.mm น.')}</TableCell>
                            </TableRow>
                          ))}
                        </TableBody>
                      </Table>
                    </TableContainer>
                  )
                    : tax == "" ? (
                      <div>
                        <TableContainer component={Paper}>
                          <Table className={classes.table} aria-label="simple table">
                            <TableHead>
                              <TableRow>
                                <TableCell align="center">ลำดับที่</TableCell>
                                <TableCell align="center">เลขที่กำกับภาษี</TableCell>
                                
                                <TableCell align="center">ชื่อผู้ชำระค่าบริการ</TableCell>
                                <TableCell align="center">ชื่อผู้ป่วย</TableCell>
                                <TableCell align="center">ประเภทการชำระ</TableCell>
                                <TableCell align="center">วันที่</TableCell>
                              </TableRow>
                            </TableHead>
                            <TableBody>

                              {dentalexpense.map((item: any) => (
                                <TableRow key={item.id}>
                                  <TableCell align="center">{item.id}</TableCell>
                              <TableCell align="center">{item.tax}</TableCell>
                              
                              <TableCell align="center">{item.name}</TableCell>
                              <TableCell align="center">{item.edges?.medicalfile?.edges?.patient?.name}</TableCell>
                              <TableCell align="center">{item.edges?.pricetypes?.id}</TableCell>
                              <TableCell align="center">{moment(item.AddedTime).format('DD/MM/YYYY HH.mm น.')}</TableCell>
                                </TableRow>
                              ))}
                            </TableBody>
                          </Table>
                        </TableContainer>

                      </div>
                    ) : null}
                </div>
              ) : null}
            </Paper>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );

}

import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import { Link as RouterLink } from 'react-router-dom';

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Avatar,
  Button,
  Link,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntMedicalfile } from '../../api/models/EntMedicalfile'; // import interface Dentist
import { EntPriceType } from '../../api/models/EntPriceType'; // import interface Nurse
import { EntNurse } from '../../api/models/EntNurse'; // import interface Patient


// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },

  menuButton: {
    marginRight: theme.spacing(100),
  },
  title: {
    flexGrow: 1,
  },

}));

interface saveDenExpen {
  Medicalfile: Number;
  PriceType: Number;
  Nurse: Number;
  Rates: Number;
  Added: String;
  Tax: String;
  Name: String;
  Phone: String;
}

const SaveDenExpen: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [dentalexpense, setDentalExpense] = React.useState<Partial<saveDenExpen>>({});
  const [medicalfiles, setMedicalfiles] = React.useState<EntMedicalfile[]>([]); //การประกาศตัวแปร โดยที่เราจะดึงมาใช้ แล้ว Ent ได้มาจากการเจน API
  const [pricetypes, setPriceTypes] = React.useState<EntPriceType[]>([]);
  const [nurses, setNurses] = React.useState<EntNurse[]>([]);

  // alert setting
  //const Swal = require('sweetalert2')
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
  });

  // set data to object medicalfile
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof dentalexpense;
    const { value } = event.target;
    setDentalExpense({ ...dentalexpense, [name]: value });
    console.log(dentalexpense);
  };
  
  const getPriceType = async () => {
    const res = await http.listPricetype({ limit: 4, offset: 0 });
    setPriceTypes(res);
  };

  const getMedicalfile = async () => {
    const res = await http.listMedicalfile({ limit: 4, offset: 0 });
    setMedicalfiles(res);
  };

  const getNurse = async () => {
    const res = await http.listNurse({ limit: 4, offset: 0 });
    setNurses(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getMedicalfile();
    getPriceType();
    getNurse();
  }, []);

  // clear input form
  function clear() {
    setDentalExpense({});
  }

  // function save data
  function save() {
    dentalexpense.Added += ":00+07:00";
    const apiUrl = 'http://localhost:8080/api/v1/dentalexpenses';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(dentalexpense),
    };

    console.log(dentalexpense); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {console.log(data.save);
        console.log(requestOptions)
        if (data.status == true) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }

  return (
    <Page theme={pageTheme.service}>
      <Header
       title="Dental System"
       subtitle="ระบบบันทึกค่ารักษา">
     </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>เลขที่กำกับภาษี</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                
                label="ระบุเลขที่กำกับภาษี" 
                variant="outlined" 
                name="Tax"
                type="string"
                value={dentalexpense.Tax || ''}
                onChange={handleChange}
                />

              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ข้อมูลการรับบริการทันตกรรม</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกข้อมูลการรับบริการทันตกรรม</InputLabel>
                <Select
                  name="Medicalfile"
                  value={dentalexpense.Medicalfile || ''} 
                  onChange={handleChange}
                >
                  {medicalfiles.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.detail}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ค่าบริการ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                
                label="ระบุค่าบริการ" 
                variant="outlined" 
                name="Rates"
                type="string"
                value={dentalexpense.Rates || ''}
                onChange={handleChange}
                />

              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ประเภทการชำระ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกประเภทการชำระ</InputLabel>
                <Select
                  name="PriceType"
                  value={dentalexpense.PriceType || ''} 
                  onChange={handleChange}
                >
                  {pricetypes.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                      {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ชื่อผู้ชำระค่าบริการ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                
                label="ระบุชื่อผู้ชำระค่าบริการ" 
                variant="outlined" 
                name="Name"
                type="string"
                value={dentalexpense.Name || ''}
                onChange={handleChange}
                />

              </FormControl>
            </Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>เบอร์ที่สามารถติดต่อได้</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                
                label="ระบุเบอร์โทรศัพท์" 
                variant="outlined" 
                name="Phone"
                type="string"
                value={dentalexpense.Phone || ''}
                onChange={handleChange}
                />

              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>วันที่ชำระค่าบริการ</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  
                  label="เลือกวันที่"
                  name="Added"
                  type="datetime-local"
                  value={dentalexpense.Added || ''} 
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                บันทึกค่ารักษา
              </Button>
              &emsp;
              <Link component={RouterLink} to="/Menu">
              <Button
                variant="contained"
                color="default"
                size="large"
                style={{ marginLeft: 5 }}
              >
                กลับ
              </Button>
              </Link>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default  SaveDenExpen;
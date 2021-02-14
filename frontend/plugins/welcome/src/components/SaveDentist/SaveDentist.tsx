
import React, { FC, useEffect} from 'react';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import Swal from 'sweetalert2'; // alert
import SaveIcon from '@material-ui/icons/Save'; // icon save
import { Link as RouterLink } from 'react-router-dom';
import { Container,} from '@material-ui/core';
import 'date-fns';
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntDegree } from '../../api/models/EntDegree'; // import interface Degree
import { EntExpert } from '../../api/models/EntExpert'; // import interface Expert
import { EntGender } from '../../api/models/EntGender'; // import interface Gender
import { 
 Content,
 Header,
 Page,
 pageTheme,
 //ContentHeader,
 Link,
 
} from '@backstage/core';

const useStyles = makeStyles((theme) => ({
 
  button: {
    margin: theme.spacing(1, 0),
  },
   formControl: {
    width: 300,
  },

  noLabel: {
    marginTop: theme.spacing(3),
  },

  textField: {
    width: '20ch',
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
}));

interface saveDentist {
 name: string;
 age: number ;
 gender: number  ;
 cardid: string;
 birthday: string ;
 experience: string ;
 tel:string;
 email:string;
 password:string;
 expert:number;
 degree:number;
}

const SaveDentist: FC<{}> = () => { 

    const classes = useStyles();
    const http = new DefaultApi();
    const [dentist, setDentist] = React.useState<Partial<saveDentist>>({});
    const [degrees, setDegrees] = React.useState<EntDegree[]>([]);
  const [genders, setGenders] = React.useState<EntGender[]>([]);
  const [experts, setExperts] = React.useState<EntExpert[]>([]);
  const [cardiderr, setCarderr] =React.useState('');
  const [telErr, setTel] = React.useState('');
  const [emailErr, setEmailerr] = React.useState('');
  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });
  const getDegrees = async () => {
    const res = await http.listDegree({ limit: 10, offset: 0 });
    setDegrees(res);
  };
  const getGenders = async () => {
    const res = await http.listGender({ limit: 10, offset: 0 });
    setGenders(res);
  };
  const getExperts = async () => {
    const res = await http.listExpert({ limit: 10, offset: 0 });
    setExperts(res);
  };
 // Lifecycle Hooks
  useEffect(() => {
    getDegrees();
    getGenders();
    getExperts();
  }, []);

  const validateCardid = (val: string) => {
    return val.match("^[0-9]{13}") && val.length <= 13;
  }

  const varlidateTel =  (val: string) => {
    return val.match("^[0]\\d{9}")&& val.length <= 10 ;
  }

  const varlidateEmail =  (email: string) => {
    return email.match("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$");
  }

  // set data to object dentist
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,) => {
    const name = event.target.name as keyof typeof dentist;
    const { value } = event.target;
    setDentist({ ...dentist, [name]: value });
    const validateValue = value.toString()
    checkPattern(name, validateValue)
    console.log(dentist);
  };

  const checkPattern  = (s: string, value: string) => {
    switch(s) {
      case 'cardid':
        validateCardid(value) ? setCarderr('') : setCarderr('กรอกเลขบัตรประชาชน 13 หลัก');
        return;
      case 'tel':
        varlidateTel(value) ? setTel('') : setTel('กรอกหมายเลขโทรศัพท์');
        return;
      case 'email':
      varlidateEmail(value) ? setEmailerr('') : setEmailerr('กรอกอีเมลให้ถูกต้อง');
        return;
      default:
        return;
    }
  }
   const checkCaseSaveError = (field: string) => {
     switch(field){
       case 'cardid':
        alertMessage("error","กรอกเลขบัตรประชาชน 13 หลัก ผิดพลาด");
         return;
         case 'tel':
          alertMessage("error","กรอกหมายเลขโทรศัพท์ ผิดพลาด");
           return;
        case 'email':
             alertMessage("error","รูปแบบอีเมลผิดพลาด กรอกอีเมลให้ถูกต้อง");
           return;
      default:
       alertMessage("error", " บันทึกข้อมูลไม่สำเร็จ");
          return;
     }
   }

  const handleNumberChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,) => {
    const name = event.target.name as keyof typeof dentist;
    const { value } = event.target;
    setDentist({ ...dentist, [name]: + value });
    console.log(dentist);
  };

  // clear input form
  function clear() {
    setDentist({});
  }
 // function save data
 function save() {
    dentist.birthday += ":00+07:00";
    const apiUrl = 'http://localhost:8080/api/v1/dentists';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(dentist),
    };
  
    console.log(dentist); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console
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
        checkCaseSaveError(data.error.Name)
      }
    });
   }
   
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
   }

   const profile = { givenName: 'Detal System' };
   return(
    <Page theme={pageTheme.service}>
    <Header
      title={` ${profile.givenName || 'to Backstage'}`}
     subtitle="ระบบบันทึกข้อมูลทันตแพทย์"  
    ></Header>
    <Content>
    <Container maxWidth="sm">
    <Grid container spacing={3}>
       <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ชื่อ - นามสกุล</div>
            </Grid>
     <Grid item xs={9}>
       <FormControl fullWidth variant="outlined" className={classes.formControl}>
              <TextField 
              label="ชื่อ - นามสกุล" 
              variant="outlined" 
              name="name"
              type="string"
              value={dentist.name || ''}
              onChange={handleChange}
              />
   
       </FormControl>
       </Grid>

     <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>อายุ</div>
            </Grid>
     <Grid item xs={9}>
            <FormControl
                  fullWidth
                  className={classes.formControl}
                  variant="outlined">
                  <TextField
                    label="อายุ"
                    name="age"
                    type="Number"
                    variant="outlined"
                    size="medium"
                    value={dentist.age|| ''}
                    onChange={handleNumberChange}
                  />
                </FormControl>
       </Grid>
    
     <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>เกิด - วันที่</div>
            </Grid>
     <Grid item xs={9}>
        <FormControl variant="outlined" className={classes.formControl}>
             <form className={classes.container} noValidate>
               <TextField
                 label="เกิด - วันที่"
                 name="birthday"
                 type="datetime-local"
                 value={dentist.birthday|| ''} 
                 className={classes.textField}
                 InputLabelProps={{
                   shrink: true,
                 }}
                 onChange={handleChange}
               />
             </form>
               
        </FormControl>
       </Grid>     
      
     <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>หมายเลขบัตรประชาชน</div>
            </Grid>
     <Grid item xs={9}>
     <FormControl  fullWidth variant="outlined" className={classes.formControl}>
              <TextField 
              error = {cardiderr ? true : false}
              label="หมายเลขบัตรประชาชน" 
              inputProps={{ maxLength: 13 }}
              variant="outlined" 
              name="cardid"
              type="string"
              helperText= {cardiderr}
              value={dentist.cardid|| ''}
              onChange={handleChange}
              />
     </FormControl>
     </Grid>


     
     <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>วุฒิการศึกษา</div>
            </Grid>
     <Grid item xs={9}>
        <FormControl variant="outlined" className={classes.formControl}>
       <InputLabel >วุฒิการศึกษา</InputLabel>
       <Select
         name="degree"
         value={dentist.degree || ''} // (undefined || '') = ''
         onChange={handleChange}
         
       >
          {degrees.map(item => {
                   return (
                     <MenuItem key={item.id} value={item.id}>
                       {item.name}
                     </MenuItem>
                   );
           })}
       </Select>
       </FormControl>
       </Grid>

       <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>เพศ</div>
            </Grid>
       <Grid item xs={9}>
   <FormControl variant="outlined" className={classes.formControl}>
       <InputLabel >เพศ</InputLabel>
       <Select
         name="gender"
         value={dentist.gender || ''} // (undefined || '') = ''
         onChange={handleChange}
         
       >
          {genders.map(item => {
                   return (
                     <MenuItem key={item.id} value={item.id}>
                       {item.name}
                     </MenuItem>
                   );
           })}
           </Select>
     </FormControl>
     </Grid>

   <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ความเชี่ยวชาญทางทันกรรม</div>
            </Grid>
   <Grid item xs={9}>
   <FormControl variant="outlined" className={classes.formControl}>
       <InputLabel >ความเชี่ยวชาญทางทันกรรม</InputLabel>
       <Select
         name="expert"
         value={dentist.expert || ''} // (undefined || '') = ''
         onChange={handleChange}
       >
          {experts.map(item => {
                   return (
                     <MenuItem key={item.id} value={item.id}>
                       {item.name}
                     </MenuItem>
                   );
           })}          
       </Select>
     </FormControl>
     </Grid>   
   
         <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>ประสบการณ์ทำงาน</div>
            </Grid>
     <Grid item xs={9}>
     <FormControl variant="outlined" className={classes.formControl}>
              <TextField 
              label="ประสบการณ์ทำงาน" 
              variant="outlined" 
              name="experience"
              multiline
              rows={4}
              type="string"
              size="medium"
              value={dentist.experience || ''}
              onChange={handleChange}
              />
       </FormControl>
    </Grid>   
   
     <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>เบอร์โทรศัพท์</div>
            </Grid>
     <Grid item xs={9}>
     <FormControl variant="outlined" className={classes.formControl}>
              <TextField 
              error = {telErr ? true : false}
              label="เบอร์โทรศัพท์" 
              inputProps={{ maxLength: 10 }}
              variant="outlined" 
              name="tel"
              type="string"
              helperText= {telErr}
              value={dentist.tel || ''}
              onChange={handleChange}
              />
     </FormControl>
    </Grid>

     <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>Email</div>
            </Grid>
     <Grid item xs={9}>
     <FormControl variant="outlined" className={classes.formControl}>
             <TextField 
              error = {emailErr ? true : false}
              label="Email" 
              variant="outlined" 
              name="email"
              type="string"
              helperText= {emailErr}
              value={dentist.email|| ''}
              onChange={handleChange}
              />
     </FormControl>
     </Grid>
     <Grid item xs={12}></Grid>   
     <Grid item xs={3}>
              <div className={classes.paper}>Password</div>
            </Grid>
         <Grid item xs={9}>
            <FormControl variant="outlined" className={classes.formControl}>
              <TextField 
              label="Password" 
              variant="outlined" 
              name="password"
              type="string"
              value={dentist.password || ''}
              onChange={handleChange}
              />
              </FormControl>
           </Grid>

            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                style={{ marginRight: 5 }}
                startIcon={<SaveIcon />}
                onClick={save}
              >
                บันทึกข้อมูล
              </Button>
              &emsp;
              <Link component={RouterLink} to="/menu">
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
export default  SaveDentist;
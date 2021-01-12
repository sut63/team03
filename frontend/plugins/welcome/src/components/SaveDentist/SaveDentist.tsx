
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
//import { Container,} from '@material-ui/core';
import 'date-fns';
import { makeStyles } from '@material-ui/core/styles';
import Grid from '@material-ui/core/Grid';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntDegree } from '../../api/models/EntDegree'; // import interface Degree
//import { EntDentist } from '../../api/models/EntDentist'; // import interface Dentist
import { EntExpert } from '../../api/models/EntExpert'; // import interface Expert
import { EntGender } from '../../api/models/EntGender'; // import interface Gender
//import { EntNurse } from '../../api/models/EntNurse'; // import interface Nurse
//import { Container } from '@material-ui/core';
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
  margin: {
    margin: theme.spacing(3),
  },
  textField: {
    width: '20ch',
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
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


  // set data to object dentist
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,) => {
    const name = event.target.name as keyof typeof dentist;
    const { value } = event.target;
    setDentist({ ...dentist, [name]: value });
    
    console.log(dentist);
  };

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
        Toast.fire({
          icon: 'error',
          title: 'บันทึกข้อมูลไม่สำเร็จ',
        });
      }
    });
  }

const profile = { givenName: 'Detal System' };
return(
    <Page theme={pageTheme.home}>
    <Header
      title={`Welcome ${profile.givenName || 'to Backstage'}`}
     subtitle="ระบบบันทึกข้อมูลทันตแพทย์"  // <dd> <ComponanceTable></ComponanceTable></dd>
    ></Header>
    <Content>
 
    <div style={{textAlign: "center"}}>
   
  
     <br></br>
     <Grid item xs={9}>
     <FormControl variant="outlined" className={classes.formControl}>
              <TextField 
              label="ชื่อ - นามสกุล" 
              variant="outlined" 
              name="name"
              type="string"
              value={dentist.name || ''}
              onChange={handleChange}
              />
         <br></br>
     </FormControl>
     <Grid item xs={12}>
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
     <br></br>
     <br></br>
     <Grid item xs={12}>
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
    <br></br>
     <br></br>

     <FormControl variant="outlined" className={classes.formControl}>
              <TextField 
              label="หมายเลขบัตรประชาชน" 
              variant="outlined" 
              name="cardid"
              type="string"
              value={dentist.cardid|| ''}
              onChange={handleChange}
              />
     </FormControl>
          
    <br></br>
     <br></br>
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
     </Grid>&emsp; &emsp;
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
     <br></br>
   <br></br>
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

         <br></br>
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
          
    <br></br>
     <br></br>
     </Grid>   
     <br></br>
    
     <br></br>
     <Grid item xs={9}>
     <FormControl variant="outlined" className={classes.formControl}>
              <TextField 
              label="เบอร์โทรศัพท์" 
              variant="outlined" 
              name="tel"
              type="string"
              value={dentist.tel || ''}
              onChange={handleChange}
              />
     </FormControl>
    </Grid>
  
     <br></br>
     <br></br>

     <Grid item xs={9}>
     <FormControl variant="outlined" className={classes.formControl}>
             
             <TextField 
              label="Email" 
              variant="outlined" 
              name="email"
              type="string"
              value={dentist.email|| ''}
              onChange={handleChange}
              />
                <br></br>
              <TextField 
              label="password" 
              variant="outlined" 
              name="password"
              type="string"
              value={dentist.password || ''}
              onChange={handleChange}
              />

     </FormControl>
          
    <br></br>
     <br></br>
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
       บันทึกข้อมูล
     </Button>&emsp;
     
     <Link component={RouterLink} to="/homepage">
             <Button
               variant="contained"
               color="default"
             >
              กลับ
             </Button>
             </Link>
     </Grid>                
    </div>
    </Content>
 </Page>
);
};
export default  SaveDentist;
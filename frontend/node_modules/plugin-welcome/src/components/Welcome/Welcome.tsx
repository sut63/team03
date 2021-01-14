import React, { FC } from 'react';
import Button from '@material-ui/core/Button';
import ExitToAppRoundedIcon from '@material-ui/icons/ExitToAppRounded';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import Link from '@material-ui/core/Link';
import { makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import ButtonGroup from '@material-ui/core/ButtonGroup';

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    justifyContent: 'center'
  },
  paper: {
    height: 330,
    width: 1250,
    backgroundImage: 'https://thumbs.dreamstime.com/z/unhealthy-vs-healthy-teeth-cartoon-comparison-illustration-vector-tooth-decay-destroyed-cute-characters-dental-care-background-142147889.jpg)',
  },
  control: {
    padding: theme.spacing(2),
  },
 
})
);

const Homepage: FC<{}> = () => {
  const classes = useStyles();
  return (
    <Page theme={pageTheme.service}>
      <Content>

        <center>
        <ContentHeader title="Welcome"> </ContentHeader>
        <ButtonGroup size="large" aria-label="small outlined button group">
        <Button><Link color="textPrimary" href="/SignInDentist" > login ในฐานะทันตแพทย์ </Link></Button>
        <Button><Link color="textPrimary" href="/SignInNurse" > login ในฐานะพยาบาล </Link></Button>
        </ButtonGroup>
        </center>

      </Content>
    </Page>
  );
};
export default Homepage;

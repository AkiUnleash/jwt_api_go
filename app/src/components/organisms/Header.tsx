import React from 'react';
import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import axios from 'axios'
import { browserHistory } from '../../history'

const useStyles = makeStyles((theme) => ({
  '@global': {
    ul: {
      margin: 0,
      padding: 0,
      listStyle: 'none',
    },
  },
  appBar: {
    borderBottom: `1px solid ${theme.palette.divider}`,
  },
  toolbar: {
    flexWrap: 'wrap',
  },
  toolbarTitle: {
    flexGrow: 1,
  },
  link: {
    margin: theme.spacing(1, 1.5),
  },
}));

const Header: React.FC = () => {

  const classes = useStyles();

  return (
    <AppBar position="static" color="default" elevation={0} className={classes.appBar}>
      <Toolbar className={classes.toolbar}>
        <Typography variant="h6" color="inherit" noWrap className={classes.toolbarTitle}>
          ひとこと日記
          </Typography>
        <Button href="#"
          color="primary"
          variant="outlined"
          className={classes.link}
          onClick={async (e) => {

            e.preventDefault()

            // ログイン処理
            await axios.post('http://localhost:8082/account/logout', {}, {
              xsrfHeaderName: 'X-CSRF-Token',
              withCredentials: true
            }).then((e) => console.log(e));

            browserHistory.push('login')

          }}>
          ログアウト
          </Button>
      </Toolbar>
    </AppBar>
  );
}

export default Header
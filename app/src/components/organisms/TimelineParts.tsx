import React from 'react';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import TimelineItem from '@material-ui/lab/TimelineItem';
import TimelineSeparator from '@material-ui/lab/TimelineSeparator';
import TimelineConnector from '@material-ui/lab/TimelineConnector';
import TimelineContent from '@material-ui/lab/TimelineContent';
import TimelineOppositeContent from '@material-ui/lab/TimelineOppositeContent';
import TimelineDot from '@material-ui/lab/TimelineDot';
import AccessibilityNewIcon from '@material-ui/icons/AccessibilityNew';
import Paper from '@material-ui/core/Paper';

const useStyles = makeStyles((theme) => ({
  '@global': {
    ul: {
      margin: 0,
      padding: 0,
      listStyle: 'none',
    },
  },
  paper: {
    padding: '6px 16px',
  },
}));

type Diary = {
  body: string
}

const TimelineParts: React.FC<Diary> = (props: Diary) => {

  const classes = useStyles();

  return (
    <TimelineItem>
      <TimelineOppositeContent>
        <Typography variant="body2" color="textSecondary">
          9:30 am
          </Typography>
      </TimelineOppositeContent>
      <TimelineSeparator>
        <TimelineDot>
          <AccessibilityNewIcon />
        </TimelineDot>
        <TimelineConnector />
      </TimelineSeparator>
      <TimelineContent>
        <Paper elevation={3} className={classes.paper}>
          <Typography>{props.body}</Typography>
        </Paper>
      </TimelineContent>
    </TimelineItem>
  );
}

export default TimelineParts
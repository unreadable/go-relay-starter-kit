import 'babel-polyfill';
import React from 'react';
import ReactDOM from 'react-dom';
import Relay from 'react-relay';
import {Router, Route, browserHistory, applyRouterMiddleware} from 'react-router';
import useRelay from 'react-router-relay';

class HelloApp extends React.Component {
  render() {
    const {hello} = this.props.greetings;
    return <h1>{hello}</h1>;
  }
}

HelloApp = Relay.createContainer(HelloApp, {
  fragments: {
    greetings: () => Relay.QL`
      fragment on Greetings {
        hello,
      }
    `,
  },
});

const HelloQueries = {
    greetings: (Component) => Relay.QL`
      query GreetingsQuery {
        greetings {
          ${Component.getFragment('greetings')},
        },
      }
    `,
  };

ReactDOM.render(
  <Router forceFetch environment={Relay.Store} render={applyRouterMiddleware(useRelay)} history={browserHistory} >
    <Route path='/' component={HelloApp} queries={HelloQueries} />
  </Router>, document.getElementById('root'))



import * as React from 'react';
import * as ReactDOM from 'react-dom';
import Hello from './components/Hello';

class App extends React.Component {
    render() {
        return (
            <div>
                <Hello />
            </div>
        );
    }
}


ReactDOM.render(<App />, document.querySelector('#app'));

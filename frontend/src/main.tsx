import * as React from 'react';
import * as ReactDOM from 'react-dom';
import Header from './components/Header';
import TranslateBox from './components/TranslateBox';

class App extends React.Component {
    render() {
        return (
            <>
                <Header>keiGo</Header>
                <main>
                    <TranslateBox />
                </main>
            </>
        );
    }
}


ReactDOM.render(<App />, document.querySelector('#app'));

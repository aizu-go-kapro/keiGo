import * as React from 'react';
import * as ReactDOM from 'react-dom';
import TranslateBox from './components/TranslateBox';

class App extends React.Component {
    render() {
        return (
            <>
                <header>keiGo</header>
                <main>
                    <TranslateBox />
                </main>
            </>
        );
    }
}


ReactDOM.render(<App />, document.querySelector('#app'));

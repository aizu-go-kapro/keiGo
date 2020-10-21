import * as React from 'react';
import * as ReactDOM from 'react-dom';
import Header from './components/Header';
import KeiGoBox from './components/KeiGoBox';

class App extends React.Component {
    render() {
        return (
            <>
                <Header>keiGo</Header>
                <main>
                    <KeiGoBox />
                </main>
            </>
        );
    }
}


ReactDOM.render(<App />, document.querySelector('#app'));

var React = require('react');
var ReactDOM = require('react-dom');
import styled from 'styled-components';

const StyledApp = styled.div`
    border: 1px solid #f00;
`;


export function App() {
    return (
        <StyledApp>
            Hello World!
        </StyledApp>
    )
}

if (document.getElementById('react_root')) {
    ReactDOM.render(<App />, document.getElementById('react_root'));
}
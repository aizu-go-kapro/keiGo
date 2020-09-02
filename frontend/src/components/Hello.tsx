import * as React from 'react';
import styled from 'styled-components';

type Props = {}

const Hello: React.SFC<Props> = props => {
    return (
        <BigTitle>Hello, World!</BigTitle>
    )
};

const BigTitle = styled.p`
    font-size: 3rem;
    font-weight: bold;
    color: red;
`;

export default Hello;

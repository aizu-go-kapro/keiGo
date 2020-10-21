import * as React from 'react';
import styled from 'styled-components';
import Color from '../const/Color';
import { media } from '../utils/ResponsiveHelper';

type Props = {}

const Header: React.SFC<Props> = props => {
    return (
        <Wrapper>
            <Logo>
                keiGo
            </Logo>
        </Wrapper>
    )
};

const Wrapper = styled.header`
    width: 100%;
    height: 66px;

    border: 1px solid ${Color.BORDER};
`;

const Logo = styled.div`
    display: flex;
    align-items: center;
    width: 81px;
    height: 100%;
    margin-left: 66px;

    font-family: sans-serif;
    font-style: italic;
    font-weight: bold;
    font-size: 32px;
    line-height: 37px;

    color: ${Color.PRIMARY};

    ${media.phone`
        margin-left: 33px;
    `}
`;

export default Header;
import * as React from "react";
import styled from "styled-components";
import Color from "../const/Color";
import { media } from "../utils/Helper";

interface IState {
  kind: "teinei" | "sonkei" | "kenjyo"
  body: string;
  convertedBody: string;
}

export default class TranslateBox extends React.Component<{}, IState> {
  constructor(props) {
    super(props);
    this.state = {
      kind: "teinei",
      body: "",
      convertedBody: ""
    };

    this.handleRadioChange = this.handleRadioChange.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.handleKeyPress = this.handleKeyPress.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleRadioChange(event) {
    this.setState({kind: event.target.id});
    console.log(event.target.id);
    this.postTranslate();
  }

  handleChange(event) {
    this.setState({body: event.target.value});
  }

  async handleKeyPress(event) {
    if(event.key === "Enter"){
      this.postTranslate();
    }
  }

  async postTranslate() {
    if(this.state.body === "") {
      return;
    } else {
      const url = `http://34.71.216.160:3000/api/v1/keigo?kind=${this.state.kind}`;
      const body = {
        "body": this.state.body
      };
      try {
        const response = await fetch(url, {
          method: "POST",
          headers: {
            "Content-Type": "application/json"
          },
          body: JSON.stringify(body)
        });
        const data = await response.json();
        console.log(data);
        this.setState({convertedBody: data.converted_body});
      } catch (err) {
        console.error(err);
      }
    }
  }

  render() {
    return (
      <Wrapper>
        <TranslateBar>
          <TranslateBarLeft>
            <LangBox>
              <label>
                原文
              </label>
            </LangBox>
          </TranslateBarLeft>
          <TranslateBarRight>
            <LangBoxRadioButton>
              <input type="radio" name="敬語" value="1" id="teinei" defaultChecked onChange={this.handleRadioChange} />
              <label htmlFor="teinei">丁寧</label>
            </LangBoxRadioButton>
            <LangBoxRadioButton>
              <input type="radio" name="敬語" value="2" id="sonkei" onChange={this.handleRadioChange} />
              <label htmlFor="sonkei">尊敬</label>
            </LangBoxRadioButton>
            <LangBoxRadioButton>
              <input type="radio" name="敬語" value="3" id="kenjyo" onChange={this.handleRadioChange} />
              <label htmlFor="kenjyo">謙譲</label>
            </LangBoxRadioButton>
          </TranslateBarRight>
        </TranslateBar>
        <TextBox>
          <input type="text" placeholder="テキストを入力してください" value={this.state.body} onChange={this.handleChange} onKeyPress={this.handleKeyPress} />
        </TextBox>
        <TextBox>
          <input type="text" readOnly placeholder="変換結果" value={this.state.convertedBody} />
        </TextBox>
      </Wrapper>
    );
  }
}

const Wrapper = styled.div`
  width: calc(100% - 2 * 66px);
  padding: 66px;

  ${media.phone`
    width: 100%;
    padding: 0px;
    padding-top: 66px;
  `}
`;

const TranslateBar = styled.div`
  display: flex;
  width: 100%;
  height: 48px;

  border-top: 1px solid ${Color.BORDER};
  border-left: 1px solid ${Color.BORDER};
  border-right: 1px solid ${Color.BORDER};
`;

const TranslateBarLeft = styled.div`
  display: flex;
  width: 50%;
  height: 100%;
`;

const TranslateBarRight = styled.div`
  display: flex;
  width: 50%;
  height: 100%;
`;

const LangBox = styled.div`
  width: 150px;
  height: 48px;

  label {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100%;
    height: 100%;

    font-style: normal;
    font-weight: bold;
    font-size: 16px;
    line-height: 19px;

    color: ${Color.PRIMARY};
  }
`;

const LangBoxRadioButton = styled(LangBox)`
  input[type=radio] {
    display: none;
  }
  input[type="radio"] + label {
    color: ${Color.GRAY};
  }
  input[type="radio"]:checked + label {
    color: ${Color.PRIMARY};
    box-shadow: inset 0px -3px 0px ${Color.PRIMARY};
  }
`;

const TextBox = styled.div`
  float: left;
  width: 50%;
  height: 200px;

  input {
    width: calc(100% - 2 * 36px);
    height: calc(100% - 2 * 36px);
    outline: none;
    border: 1px solid ${Color.BORDER};

    padding: 36px;

    font-weight: 500;
    font-size: 24px;
    line-height: 28px;

    color: ${Color.BLACK};
  }

  ${media.phone`
    float: none;
    width: 100%;
  `}
`;
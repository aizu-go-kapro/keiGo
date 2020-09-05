import * as React from "react";

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
    this.setState({kind: event.target.value});
    console.log(event.target.value);
  }

  handleChange(event) {
    this.setState({body: event.target.value});
  }

  async handleKeyPress(event) {
    if(event.key === "Enter"){
      const url = `http://34.71.216.160:3000/api/v1/keigo?kind=${this.state.kind}`;
      const body = {
        "body": event.target.value
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

  handleSubmit(event) {
    // ToDo: something
  }

  render() {
    return (
      <>
        <input type="radio" name="敬語" value="teinei" defaultChecked onChange={this.handleRadioChange} />丁寧
        <input type="radio" name="敬語" value="sonkei" onChange={this.handleRadioChange} />尊敬
        <input type="radio" name="敬語" value="kenjyo" onChange={this.handleRadioChange} />謙譲
        <form onSubmit={this.handleSubmit}>
          <label>
            原文:
            <textarea placeholder="テキストを入力してください" value={this.state.body} onChange={this.handleChange} onKeyPress={this.handleKeyPress} />
          </label>
          <input type="submit" value="Submit" />
        </form>
        <textarea readOnly placeholder="変換結果" value={this.state.convertedBody} />
      </>
    );
  }
}
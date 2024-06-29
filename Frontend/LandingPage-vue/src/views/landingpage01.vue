<template>
  <div class="landingpage01-container">
    <div class="landingpage01-landingpage01">
      <div class="landingpage01-frame1321316142">
        <div class="landingpage01-copy">
          <span class="landingpage01-text Title">
            <span>
              Empower Learning
              <span v-html="rawh18b"></span>
            </span>
            <br />
            <span>One connection at a time</span>
          </span>
          <span class="landingpage01-text04 Subheading">
            <span>The search for your perfect teacher starts now</span>
          </span>
          <button class="landingpage01-button">
            <span class="landingpage01-text06 Smalltext">
              <span>Begin Searching</span>
            </span>
          </button>
        </div>
      </div>
      <div class="landingpage01-image-dragintofill"></div>
      <div class="landingpage01-heading">
        <span class="landingpage01-text08 Heading"><span>Top Tutors</span></span>
        <span class="landingpage01-text10 Subheading">
          <span>See what others have to say about our tutors</span>
        </span>
      </div>
      <div class="landingpage01-cards">
        <div class="landingpage01-customer-quote">
          <span class="landingpage01-text12 Bodytext">
            <span>“A terrific piece of praise”</span>
          </span>
          <div class="landingpage01-avatar">
            <img
              src="/external/avatar1189-7174-200w.png"
              alt="Avatar1189"
              class="landingpage01-avatar1"
            />
            <div class="landingpage01-frame2610301">
              <span class="landingpage01-text14 Smalltext">
                <span>Name</span>
              </span>
              <span class="landingpage01-text16 Smalltext">
                <span>Description</span>
              </span>
            </div>
          </div>
        </div>
        <div class="landingpage01-customer-quote1">
          <span class="landingpage01-text18 Bodytext">
            <span>“A fantastic bit of feedback”</span>
          </span>
          <div class="landingpage01-avatar2">
            <img
              src="/external/avatar1196-1uuf-200h.png"
              alt="Avatar1196"
              class="landingpage01-avatar3"
            />
            <div class="landingpage01-frame26103011">
              <span class="landingpage01-text20 Smalltext">
                <span>Name</span>
              </span>
              <span class="landingpage01-text22 Smalltext">
                <span>Description</span>
              </span>
            </div>
          </div>
        </div>
        <div class="landingpage01-customer-quote2">
          <span class="landingpage01-text24 Bodytext">
            <span>“A genuinely glowing review”</span>
          </span>
          <div class="landingpage01-avatar4">
            <img
              src="/external/avatar1203-0ynt-200h.png"
              alt="Avatar1203"
              class="landingpage01-avatar5"
            />
            <div class="landingpage01-frame26103012">
              <span class="landingpage01-text26 Smalltext">
                <span>Name</span>
              </span>
              <span class="landingpage01-text28 Smalltext">
                <span>Description</span>
              </span>
            </div>
          </div>
        </div>
      </div>
      <div class="landingpage01-navigation">
        <div class="landingpage01-items">
          <span class="landingpage01-text30 Bodytext"><span>Teachers</span></span>
          <span class="landingpage01-text32 Bodytext">
            <span>Contact Us</span>
          </span>
          <span class="landingpage01-text34 Bodytext"><span>FAQ</span></span>
          <button class="landingpage01-button1">
            <span class="landingpage01-text36 Smalltext">
              <span>Sign In</span>
            </span>
          </button>
        </div>
        <span class="landingpage01-text38"><span>ConnectX</span></span>
      </div>
      <div class="landingpage01-navigation-footer">
        <span class="landingpage01-text40 Subheading"><span>ConnectX</span></span>
        <div class="landingpage01-social-icons">
          <button class="landingpage01-buttons-icon">
            <div class="landingpage01-icon">
              <img
                src="/external/icon1235-3n7o.svg"
                alt="Icon1235"
                class="landingpage01-icon1"
              />
            </div>
          </button>
          <button class="landingpage01-buttons-icon1">
            <div class="landingpage01-icon2">
              <img
                src="/external/icon1240-nmc6.svg"
                alt="Icon1240"
                class="landingpage01-icon3"
              />
            </div>
          </button>
          <button class="landingpage01-buttons-icon2">
            <div class="landingpage01-icon4">
              <img
                src="/external/icon1250-7n2o.svg"
                alt="Icon1250"
                class="landingpage01-icon5"
              />
            </div>
          </button>
        </div>
        <img
          src="/external/divider1253-i5a.svg"
          alt="Divider1253"
          class="landingpage01-divider"
        />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Landingpage01',
  data() {
    return {
      rawh18b: ' ',
    }
  },
  metaInfo: {
    title: 'exported project',
  },
}
</script>

<style scoped>
.landingpage01-container {
  width: 100%;
  display: flex;
  overflow: auto;
  min-height: 100vh;
  align-items: center;
  flex-direction: column;
}
.landingpage01-landingpage01 {
  width: 100%;
  height: 2077px;
  display: flex;
  overflow: hidden;
  position: relative;
  align-items: flex-start;
  flex-shrink: 0;
  background-color: rgba(255, 255, 255, 1);
}
.landingpage01-frame1321316142 {
  gap: 10px;
  top: 237px;
  left: 291px;
  width: 857px;
  display: flex;
  padding: 10px;
  position: absolute;
  align-items: flex-start;
  flex-direction: column;
}
.landingpage01-copy {
  gap: 24px;
  width: 837px;
  display: flex;
  align-items: center;
  flex-shrink: 0;
  flex-direction: column;
}
.landingpage01-text {
  color: rgba(0, 0, 0, 1);
  height: auto;
  align-self: stretch;
  text-align: center;
  line-height: normal;
}
.landingpage01-text04 {
  color: rgba(130, 130, 130, 1);
  height: auto;
  align-self: stretch;
  text-align: center;
  line-height: 150%;
}
.landingpage01-button {
  gap: 8px;
  display: flex;
  padding: 14px 24px;
  box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.05000000074505806) ;
  align-items: center;
  border-radius: 8px;
  justify-content: center;
  background-color: rgba(0, 0, 0, 1);
}
.landingpage01-text06 {
  color: rgba(255, 255, 255, 1);
  height: auto;
  text-align: left;
  line-height: 150%;
}
.landingpage01-image-dragintofill {
  top: 607px;
  left: 70px;
  width: 1280px;
  height: 640px;
  display: flex;
  overflow: hidden;
  position: absolute;
  align-items: flex-start;
  flex-shrink: 0;
  border-radius: 8px;
  background-size: cover;
  background-image: url("/external/tq_lze-kuk5b1-iifh-700h.png");
}
.landingpage01-heading {
  gap: 8px;
  top: 1381px;
  left: 80px;
  width: 733px;
  display: flex;
  position: absolute;
  align-items: flex-start;
  flex-shrink: 0;
  flex-direction: column;
}
.landingpage01-text08 {
  color: rgba(0, 0, 0, 1);
  height: auto;
  align-self: stretch;
  text-align: left;
  line-height: 110.00000238418579%;
}
.landingpage01-text10 {
  color: rgba(130, 130, 130, 1);
  width: 625px;
  height: auto;
  text-align: left;
  line-height: 150%;
}
.landingpage01-cards {
  gap: 32px;
  top: 1509px;
  left: 80px;
  width: 1280px;
  display: flex;
  position: absolute;
  align-items: flex-start;
  flex-shrink: 0;
}
.landingpage01-customer-quote {
  gap: 48px;
  width: 405.3333435058594px;
  display: flex;
  padding: 32px;
  flex-grow: 1;
  box-shadow: -4px 8px 20px 0px rgba(0, 0, 0, 0.10000000149011612) ;
  align-items: flex-start;
  flex-shrink: 0;
  border-color: rgba(247, 247, 247, 1);
  border-style: solid;
  border-width: 1px;
  border-radius: 12px;
  flex-direction: column;
  background-color: rgba(255, 255, 255, 1);
}
.landingpage01-text12 {
  color: rgba(0, 0, 0, 1);
  height: auto;
  align-self: stretch;
  text-align: left;
  line-height: 150%;
}
.landingpage01-avatar {
  gap: 16px;
  display: flex;
  align-items: center;
}
.landingpage01-avatar1 {
  width: 45px;
  height: 45px;
}
.landingpage01-frame2610301 {
  gap: 2px;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
}
.landingpage01-text14 {
  color: rgba(0, 0, 0, 1);
  height: auto;
  text-align: left;
  line-height: 150%;
}
.landingpage01-text16 {
  color: rgba(130, 130, 130, 1);
  height: auto;
  text-align: left;
  line-height: 150%;
}
.landingpage01-customer-quote1 {
  gap: 48px;
  display: flex;
  padding: 32px;
  flex-grow: 1;
  align-self: stretch;
  box-shadow: -4px 8px 20px 0px rgba(0, 0, 0, 0.10000000149011612) ;
  align-items: flex-start;
  border-color: rgba(247, 247, 247, 1);
  border-style: solid;
  border-width: 1px;
  border-radius: 12px;
  flex-direction: column;
  background-color: rgba(255, 255, 255, 1);
}
.landingpage01-text18 {
  color: rgba(0, 0, 0, 1);
  height: auto;
  align-self: stretch;
  text-align: left;
  line-height: 150%;
}
.landingpage01-avatar2 {
  gap: 16px;
  display: flex;
  align-items: center;
}
.landingpage01-avatar3 {
  width: 45px;
  height: 45px;
}
.landingpage01-frame26103011 {
  gap: 2px;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
}
.landingpage01-text20 {
  color: rgba(0, 0, 0, 1);
  height: auto;
  text-align: left;
  line-height: 150%;
}
.landingpage01-text22 {
  color: rgba(130, 130, 130, 1);
  height: auto;
  text-align: left;
  line-height: 150%;
}
.landingpage01-customer-quote2 {
  gap: 48px;
  display: flex;
  padding: 32px;
  flex-grow: 1;
  align-self: stretch;
  box-shadow: -4px 8px 20px 0px rgba(0, 0, 0, 0.10000000149011612) ;
  align-items: flex-start;
  border-color: rgba(247, 247, 247, 1);
  border-style: solid;
  border-width: 1px;
  border-radius: 12px;
  flex-direction: column;
  background-color: rgba(255, 255, 255, 1);
}
.landingpage01-text24 {
  color: rgba(0, 0, 0, 1);
  height: auto;
  align-self: stretch;
  text-align: left;
  line-height: 150%;
}
.landingpage01-avatar4 {
  gap: 16px;
  display: flex;
  align-items: center;
}
.landingpage01-avatar5 {
  width: 45px;
  height: 45px;
}
.landingpage01-frame26103012 {
  gap: 2px;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
}
.landingpage01-text26 {
  color: rgba(0, 0, 0, 1);
  height: auto;
  text-align: left;
  line-height: 150%;
}
.landingpage01-text28 {
  color: rgba(130, 130, 130, 1);
  height: auto;
  text-align: left;
  line-height: 150%;
}
.landingpage01-navigation {
  top: 0px;
  left: 0px;
  width: 1440px;
  height: 164px;
  display: flex;
  overflow: hidden;
  position: absolute;
  align-items: flex-start;
  flex-shrink: 0;
  background-color: rgba(255, 255, 255, 1);
}
.landingpage01-items {
  gap: 48px;
  top: 56px;
  left: 880px;
  width: 480px;
  display: flex;
  position: absolute;
  align-items: center;
  justify-content: flex-end;
}
.landingpage01-text30 {
  color: rgba(0, 0, 0, 1);
  height: auto;
  text-align: left;
  line-height: 150%;
}
.landingpage01-text32 {
  color: rgba(0, 0, 0, 1);
  height: auto;
  text-align: left;
  line-height: 150%;
}
.landingpage01-text34 {
  color: rgba(0, 0, 0, 1);
  height: auto;
  text-align: left;
  line-height: 150%;
}
.landingpage01-button1 {
  gap: 8px;
  display: flex;
  padding: 14px 24px;
  box-shadow: 0px 1px 2px 0px rgba(0, 0, 0, 0.05000000074505806) ;
  align-items: center;
  border-radius: 8px;
  justify-content: center;
  background-color: rgba(0, 0, 0, 1);
}
.landingpage01-text36 {
  color: rgba(255, 255, 255, 1);
  height: auto;
  text-align: left;
  line-height: 150%;
}
.landingpage01-text38 {
  top: 52px;
  left: 80px;
  color: rgba(0, 0, 0, 1);
  height: auto;
  position: absolute;
  font-size: 40px;
  font-style: Medium;
  text-align: left;
  font-family: Inter;
  font-weight: 500;
  line-height: 150%;
  font-stretch: normal;
  text-decoration: none;
}
.landingpage01-navigation-footer {
  top: 1852px;
  left: -1px;
  width: 1440px;
  height: 239px;
  display: flex;
  overflow: hidden;
  position: absolute;
  align-items: flex-start;
  flex-shrink: 0;
  background-color: rgba(255, 255, 255, 1);
}
.landingpage01-text40 {
  top: 86px;
  left: 664px;
  color: rgba(0, 0, 0, 1);
  height: auto;
  position: absolute;
  text-align: left;
  line-height: 150%;
}
.landingpage01-social-icons {
  gap: 8px;
  top: 132px;
  left: 652px;
  width: 136px;
  display: flex;
  position: absolute;
  align-items: flex-start;
}
.landingpage01-buttons-icon {
  gap: 8px;
  width: 40px;
  height: 40px;
  display: flex;
  padding: 8px;
  position: relative;
  align-items: center;
  border-radius: 4px;
  justify-content: center;
}
.landingpage01-icon {
  top: 8px;
  left: 8px;
  width: 24px;
  height: 24px;
  display: flex;
  position: absolute;
  align-items: flex-start;
  flex-shrink: 0;
}
.landingpage01-icon1 {
  top: 2.0399999618530273px;
  left: 2px;
  width: 20px;
  height: 20px;
  position: absolute;
}
.landingpage01-buttons-icon1 {
  gap: 8px;
  width: 40px;
  height: 40px;
  display: flex;
  padding: 8px;
  position: relative;
  align-items: center;
  border-radius: 4px;
  justify-content: center;
}
.landingpage01-icon2 {
  top: 8px;
  left: 8px;
  width: 24px;
  height: 24px;
  display: flex;
  position: absolute;
  align-items: flex-start;
  flex-shrink: 0;
}
.landingpage01-icon3 {
  top: 3px;
  left: 3px;
  width: 18px;
  height: 18px;
  position: absolute;
}
.landingpage01-buttons-icon2 {
  gap: 8px;
  width: 40px;
  height: 40px;
  display: flex;
  padding: 8px;
  position: relative;
  align-items: center;
  border-radius: 4px;
  justify-content: center;
}
.landingpage01-icon4 {
  top: 8px;
  left: 8px;
  width: 24px;
  height: 24px;
  display: flex;
  position: absolute;
  align-items: flex-start;
  flex-shrink: 0;
}
.landingpage01-icon5 {
  top: 2px;
  left: 2px;
  width: 20px;
  height: 20px;
  position: absolute;
}
.landingpage01-divider {
  top: 1px;
  left: 80px;
  width: 1280px;
  height: 1px;
  position: absolute;
  align-self: stretch;
}
</style>

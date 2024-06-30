<template>
  <div class="landingpage01-container">
    <div class="landingpage01-landingpage01">
      <div class="landingpage01-navigation">
        <span class="landingpage01-text38"><span>ConnectX</span></span>
        <div style="flex-grow: 1">&nbsp;</div>
        <div class="landingpage01-items">
          <span class="landingpage01-text30 Bodytext" @click="$router.push('/teacher')">
            <span>Teachers</span>
          </span>
          <span class="landingpage01-text32 Bodytext">
            <span>Contact Us</span>
          </span>
          <span class="landingpage01-text34 Bodytext"><span>FAQ</span></span>
          <button class="landingpage01-button1" @click="$router.push('/auth/login')">
            <span class="landingpage01-text36 Smalltext">
              <span>Sign In</span>
            </span>
          </button>
        </div>
      </div>
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
      <div class="landingpage01-navigation-footer">
        <img
            src="/external/divider1253-i5a.svg"
            alt="Divider1253"
            class="landingpage01-divider"
        />
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
      </div>
    </div>
  </div>
</template>

<script>
import button from "primevue/button";

export default {
  name: 'LandingPage',
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
  background-color: rgba(255, 255, 255, 1);
}
.landingpage01-landingpage01 {
  width: 1280px;
  height: 100%;
  display: flex;
  align-items: center;
  flex-direction: column;
  background-color: rgba(255, 255, 255, 1);
}
.landingpage01-frame1321316142 {
  gap: 10px;
  top: 237px;
  left: 291px;
  width: 857px;
  display: flex;
  padding: 10px;
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
  align-items: flex-start;
  flex-shrink: 0;
  border-radius: 8px;
  background-size: cover;
  background-image: url("/external/tq_lze-kuk5b1-iifh-700h.png");
  margin-top: 50px;
}
.landingpage01-heading {
  gap: 8px;
  top: 1381px;
  left: 80px;
  width: 733px;
  display: flex;
  align-self: flex-start;
  flex-shrink: 0;
  flex-direction: column;
  margin-top: 100px;
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
  align-items: flex-start;
  flex-shrink: 0;
  margin-top: 20px;
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
  height: 164px;
  width: 1280px;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: rgba(255, 255, 255, 1);
  margin-top: 300px;
  margin-bottom: 50px;
}
.landingpage01-items {
  gap: 48px;
  top: 56px;
  left: 880px;
  width: 480px;
  display: flex;
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
  color: rgba(0, 0, 0, 1);
  height: auto;
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
  justify-content: center;
  align-items: center;
  flex-direction: column;
  flex-shrink: 0;
  background-color: rgba(255, 255, 255, 1);
  margin-top: 100px;
}
.landingpage01-text40 {
  top: 86px;
  left: 664px;
  color: rgba(0, 0, 0, 1);
  height: auto;
  line-height: 150%;
  margin-top: 50px;
}
.landingpage01-social-icons {
  gap: 8px;
  top: 132px;
  left: 652px;
  width: 136px;
  display: flex;
  align-items: flex-start;
}
.landingpage01-buttons-icon {
  gap: 8px;
  width: 40px;
  height: 40px;
  display: flex;
  padding: 8px;
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
  align-items: flex-start;
  flex-shrink: 0;
}
.landingpage01-icon1 {
  top: 2.0399999618530273px;
  left: 2px;
  width: 20px;
  height: 20px;
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
  align-items: flex-start;
  flex-shrink: 0;
}
.landingpage01-icon3 {
  top: 3px;
  left: 3px;
  width: 18px;
  height: 18px;
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
  align-items: flex-start;
  flex-shrink: 0;
}
.landingpage01-icon5 {
  top: 2px;
  left: 2px;
  width: 20px;
  height: 20px;
}
.landingpage01-divider {
  top: 1px;
  left: 80px;
  width: 1280px;
  height: 1px;
  align-self: stretch;
}

:root {
  --dl-size-size-large: 144px;
  --dl-size-size-small: 48px;
  --dl-size-size-medium: 96px;
  --dl-size-size-xlarge: 192px;
  --dl-size-size-xsmall: 16px;
  --dl-space-space-unit: 16px;
  --dl-size-size-xxlarge: 288px;
  --dl-size-size-maxwidth: 1400px;
  --dl-color-theme-accent1: #FFFFFF;
  --dl-color-theme-accent2: #F5D1B0;
  --dl-radius-radius-round: 50%;
  --dl-color-theme-primary1: #BF4408;
  --dl-color-theme-primary2: #E65103;
  --dl-space-space-halfunit: 8px;
  --dl-space-space-sixunits: 96px;
  --dl-space-space-twounits: 32px;
  --dl-radius-radius-radius2: 2px;
  --dl-radius-radius-radius4: 4px;
  --dl-radius-radius-radius8: 8px;
  --dl-space-space-fiveunits: 80px;
  --dl-space-space-fourunits: 64px;
  --dl-color-theme-secondary1: #FFFFFF;
  --dl-color-theme-secondary2: #FBF1EB;
  --dl-space-space-threeunits: 48px;
  --dl-color-theme-neutral-dark: #191818;
  --dl-radius-radius-cardradius: 8px;
  --dl-color-theme-neutral-light: #FBFAF9;
  --dl-radius-radius-imageradius: 8px;
  --dl-radius-radius-inputradius: 24px;
  --dl-radius-radius-buttonradius: 24px;
  --dl-space-space-oneandhalfunits: 24px;
}
.button {
  color: var(--dl-color-theme-neutral-dark);
  display: inline-block;
  padding: 0.5rem 1rem;
  border-color: var(--dl-color-theme-neutral-dark);
  border-width: 1px;
  border-radius: 4px;
  background-color: var(--dl-color-theme-neutral-light);
}
.input {
  color: var(--dl-color-theme-neutral-dark);
  cursor: auto;
  padding: 0.5rem 1rem;
  border-color: var(--dl-color-theme-neutral-dark);
  border-width: 1px;
  border-radius: 4px;
  background-color: var(--dl-color-theme-neutral-light);
}
.textarea {
  color: var(--dl-color-theme-neutral-dark);
  cursor: auto;
  padding: 0.5rem;
  border-color: var(--dl-color-theme-neutral-dark);
  border-width: 1px;
  border-radius: 4px;
  background-color: var(--dl-color-theme-neutral-light);
}
.list {
  width: 100%;
  margin: 1em 0px 1em 0px;
  display: block;
  padding: 0px 0px 0px 1.5rem;
  list-style-type: none;
  list-style-position: outside;
}
.list-item {
  display: list-item;
}
.teleport-show {
  display: flex !important;
  transform: none !important;
}
.thq-input {
  color: var(--dl-color-theme-neutral-dark);
  cursor: auto;
  outline: none;
  padding: 0.5rem 1rem;
  align-self: stretch;
  text-align: center;
  border-color: var(--dl-color-theme-neutral-dark);
  border-width: 1px;
  border-radius: var(--dl-radius-radius-inputradius);
  background-color: var(--dl-color-theme-neutral-light);
}
.thq-input:focus {
  outline: 1px solid var(--dl-color-theme-primary1);
}
.thq-button-filled {
  gap: var(--dl-space-space-halfunit);
  fill: var(--dl-color-theme-secondary1);
  color: var(--dl-color-theme-secondary1);
  cursor: pointer;
  display: flex;
  transition: 0.3s;
  align-items: center;
  font-weight: bold;
  padding-top: var(--dl-space-space-halfunit);
  white-space: nowrap;
  border-color: var(--dl-color-theme-primary1);
  border-width: 1px;
  padding-left: var(--dl-space-space-oneandhalfunits);
  border-radius: var(--dl-radius-radius-buttonradius);
  padding-right: var(--dl-space-space-oneandhalfunits);
  padding-bottom: var(--dl-space-space-halfunit);
  justify-content: center;
  background-color: var(--dl-color-theme-primary1);
}
.thq-button-filled:hover {
  fill: var(--dl-color-theme-secondary2);
  color: var(--dl-color-theme-secondary2);
  border-color: var(--dl-color-theme-primary2);
  background-color: var(--dl-color-theme-primary2);
}
.thq-button-outline {
  gap: var(--dl-space-space-halfunit);
  fill: var(--dl-color-theme-primary1);
  color: var(--dl-color-theme-primary1);
  border: 1px solid;
  cursor: pointer;
  display: flex;
  transition: 0.3s;
  align-items: center;
  font-weight: bold;
  padding-top: var(--dl-space-space-halfunit);
  white-space: nowrap;
  border-color: var(--dl-color-theme-primary1);
  padding-left: var(--dl-space-space-oneandhalfunits);
  border-radius: var(--dl-radius-radius-buttonradius);
  padding-right: var(--dl-space-space-oneandhalfunits);
  padding-bottom: var(--dl-space-space-halfunit);
  justify-content: center;
}
.thq-button-outline:hover {
  fill: var(--dl-color-theme-secondary2);
  color: var(--dl-color-theme-secondary2);
  border-color: var(--dl-color-theme-primary2);
  background-color: var(--dl-color-theme-primary2);
}
.thq-button-flat {
  gap: var(--dl-space-space-halfunit);
  fill: var(--dl-color-theme-primary1);
  color: var(--dl-color-theme-primary1);
  cursor: pointer;
  display: flex;
  transition: 0.3s;
  align-items: center;
  font-weight: bold;
  padding-top: var(--dl-space-space-halfunit);
  white-space: nowrap;
  border-color: transparent;
  border-width: 1px;
  padding-left: var(--dl-space-space-oneandhalfunits);
  border-radius: var(--dl-radius-radius-buttonradius);
  padding-right: var(--dl-space-space-oneandhalfunits);
  padding-bottom: var(--dl-space-space-halfunit);
  justify-content: center;
}
.thq-button-flat:hover {
  fill: var(--dl-color-theme-secondary1);
  color: var(--dl-color-theme-secondary1);
  border-color: var(--dl-color-theme-primary2);
  background-color: var(--dl-color-theme-primary2);
}
.thq-heading-1 {
  font-size: 48px;
  font-family: STIX Two Text;
  font-weight: 700;
  line-height: 1.5;
}
.thq-heading-2 {
  font-size: 35px;
  font-family: STIX Two Text;
  font-weight: 600;
  line-height: 1.5;
}
.thq-heading-3 {
  font-size: 26px;
  font-family: STIX Two Text;
  font-weight: 600;
  line-height: 1.5;
}
.thq-body-large {
  font-size: 18px;
  font-family: Noto Sans;
  line-height: 1.5;
}
.thq-body-small {
  font-size: 16px;
  font-family: Noto Sans;
  line-height: 1.5;
}
.thq-team-image-round {
  width: 80px;
  height: 80px;
  object-fit: cover;
  border-radius: 50%;
}
.thq-section-padding {
  width: 100%;
  display: flex;
  padding: var(--dl-space-space-fiveunits);
  position: relative;
  align-items: center;
  flex-direction: column;
}
.thq-section-max-width {
  width: 100%;
  max-width: var(--dl-size-size-maxwidth);
}
.thq-img-ratio-1-1 {
  width: 100%;
  object-fit: cover;
  aspect-ratio: 1/1;
  border-radius: var(--dl-radius-radius-imageradius);
}
.thq-img-ratio-16-9 {
  width: 100%;
  object-fit: cover;
  aspect-ratio: 16/9;
  border-radius: var(--dl-radius-radius-imageradius);
}
.thq-img-ratio-4-3 {
  width: 100%;
  object-fit: cover;
  aspect-ratio: 4/3;
  border-radius: var(--dl-radius-radius-imageradius);
}
.thq-img-ratio-4-6 {
  width: 100%;
  object-fit: cover;
  aspect-ratio: 4/6;
  border-radius: var(--dl-radius-radius-imageradius);
}
.thq-img-round {
  width: 100%;
  border-radius: var(--dl-radius-radius-round);
}
.thq-flex-column {
  gap: var(--dl-space-space-twounits);
  display: flex;
  overflow: hidden;
  position: relative;
  align-items: center;
  flex-direction: column;
}
.thq-flex-row {
  gap: var(--dl-space-space-twounits);
  display: flex;
  overflow: hidden;
  position: relative;
  align-items: center;
}
.thq-grid-6 {
  display: grid;
  grid-gap: var(--dl-space-space-twounits);
  grid-template-columns: 1fr 1fr 1fr 1fr 1fr 1fr;
}
.thq-grid-5 {
  display: grid;
  grid-gap: var(--dl-space-space-twounits);
  grid-template-columns: 1fr 1fr 1fr 1fr 1fr;
}
.thq-card {
  gap: var(--dl-space-space-oneandhalfunits);
  display: flex;
  padding: var(--dl-space-space-twounits);
  align-items: stretch;
  border-radius: var(--dl-radius-radius-cardradius);
  flex-direction: column;
}
.thq-box-shadow {
  box-shadow: 0px 0px 5px -2px var(--dl-color-theme-neutral-dark);
}
.thq-grid-3 {
  display: grid;
  grid-gap: var(--dl-space-space-twounits);
  grid-template-columns: 1fr 1fr 1fr;
}
.thq-grid-4 {
  display: grid;
  grid-gap: var(--dl-space-space-twounits);
  grid-template-columns: 1fr 1fr 1fr 1fr;
}
.thq-grid-2 {
  width: 100%;
  display: grid;
  grid-gap: var(--dl-space-space-twounits);
  grid-template-columns: 1fr 1fr;
}
.thq-checkbox {
  width: var(--dl-size-size-xsmall);
  height: var(--dl-size-size-xsmall);
}
.thq-select {
  cursor: pointer;
  appearance: none;
  padding-top: var(--dl-space-space-halfunit);
  padding-left: var(--dl-space-space-unit);
  border-radius: var(--dl-radius-radius-inputradius);
  padding-right: var(--dl-space-space-twounits);
  padding-bottom: var(--dl-space-space-halfunit);
  background-color: var(--dl-color-theme-neutral-light);
  background-image: url('data:image/svg+xml;charset=US-ASCII,%3Csvg width%3D%2220%22 height%3D%2220%22 xmlns%3D%22http%3A//www.w3.org/2000/svg%22 viewBox%3D%220 0 20 20%22 fill%3D%22%23000%22%3E%3Cpath d%3D%22M4.293 7.293a1 1 0 011.414 0L10 11.586l4.293-4.293a1 1 0 111.414 1.414l-5 5a1 1 0 01-1.414 0l-5-5a1 1 0 010-1.414z%22/%3E%3C/svg%3E');
  background-repeat: no-repeat;
  background-position: right 8px center;
}
.thq-divider-horizontal {
  width: 100%;
  height: 1px;
  background-color: var(--dl-color-theme-neutral-dark);
}
.thq-icon-small {
  width: 24px;
  height: 24px;
}
.thq-button-icon {
  fill: var(--dl-color-theme-secondary1);
  padding: 3px;
  transition: 0.3s;
  border-radius: var(--dl-radius-radius-round);
}
.thq-button-icon:hover {
  fill: var(--dl-color-theme-secondary2);
}
.thq-icon-medium {
  width: var(--dl-size-size-small);
  height: var(--dl-size-size-small);
}
.thq-icon-x-small {
  width: var(--dl-size-size-xsmall);
  height: var(--dl-size-size-xsmall);
}
.thq-link {
  cursor: pointer;
  display: inline-block;
  overflow: hidden;
  background: linear-gradient(to right, var(--dl-color-theme-primary1) 50%, var(--dl-color-theme-neutral-dark) 50%);
  transition: background-position 300ms ease;
  font-weight: 600;
  background-clip: text;
  background-size: 200% 100%;
  background-position: 100%;
  -webkit-text-fill-color: transparent;
}
.thq-link:hover {
  background-position: 0 100%;
}
.thq-grid-auto-300 {
  display: grid;
  grid-gap: var(--dl-space-space-oneandhalfunits);
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
}
.thq-animated-group-vertical-reverse {
  gap: var(--dl-space-space-unit);
  width: 100%;
  display: flex;
  animation: scroll-y 20s linear infinite;
  align-items: flex-start;
  flex-direction: column;
  justify-content: space-around;
  animation-direction: reverse;
}
.thq-animated-group-horizontal-reverse {
  gap: var(--dl-space-space-unit);
  display: flex;
  animation: scroll-x 20s linear infinite;
  min-width: 100%;
  align-items: center;
  flex-shrink: 0;
  justify-content: space-around;
  animation-direction: reverse;
}
.thq-animated-group-vertical {
  gap: var(--dl-space-space-unit);
  width: 100%;
  display: flex;
  animation: scroll-y 20s linear infinite;
  align-items: flex-start;
  flex-direction: column;
  justify-content: space-around;
}
.thq-animated-group-horizontal {
  gap: var(--dl-space-space-unit);
  display: flex;
  animation: scroll-x 20s linear infinite;
  min-width: 100%;
  align-items: center;
  flex-shrink: 0;
  justify-content: space-around;
}
.thq-animated-group-container-vertical {
  gap: var(--dl-space-space-unit);
  display: flex;
  overflow: hidden;
  flex-direction: column;
}
.thq-animated-group-container-horizontal {
  gap: var(--dl-space-space-unit);
  display: flex;
  overflow: hidden;
}
.thq-mask-image-vertical {
  mask-image: linear-gradient(to bottom, transparent, black 1%, black 99%, transparent);
}
.thq-mask-image-horizontal {
  mask-image: linear-gradient(to right, transparent, black 1%, black 99%, transparent);
}
.thq-img-scale {
  transition: 0.3s;
}
.thq-img-scale:hover {
  scale: 1.05;
}
.thq-animated-card-bg-1 {
  width: 100%;
  height: 100%;
  transition: transform 0.3s;
  border-radius: var(--dl-radius-radius-cardradius);
  background-color: var(--dl-color-theme-accent1);
}
.thq-animated-card-bg-2 {
  transition: transform 0.3s;
  border-radius: var(--dl-radius-radius-cardradius);
  background-color: var(--dl-color-theme-accent2);
}
.thq-button-animated {
  outline: none;
  z-index: 1;
  overflow: hidden;
  position: relative;
  border-width: 2px;
}
.thq-input::placeholder {
  text-align: center;
  vertical-align: middle;
}
.thq-animated-group-container-vertical:hover div {
  animation-play-state: paused;
}
.thq-animated-group-container-horizontal:hover div {
  animation-play-state: paused;
}
.thq-animated-card-bg-2:has([data-animated="true"]:hover) {
  transform: translate3d(0px, 0px, 0px) scale3d(1, 1, 1) rotateX(0deg) rotateY(0deg) rotateZ(3deg) skew(0deg, 0deg);
}
.thq-animated-card-bg-1:has([data-animated="true"]:hover) {
  transform: translate3d(0px, 0px, 0px) scale3d(1, 1, 1) rotateX(0deg) rotateY(0deg) rotateZ(-6deg) skew(0deg, 0deg);
}
.thq-button-animated:before {
  top: 0;
  left: -20%;
  color: var(--dl-color-theme-neutral-light);
  width: 200%;
  height: 101%;
  content: "";
  z-index: 1;
  position: absolute;
  transform: scaleX(0);
  transition: transform 0.5s;
  border-radius: var(--dl-radius-radius-buttonradius);
  background-color: var(--dl-color-theme-neutral-dark);
  transform-origin: 0 0;
  transition-timing-function: cubic-bezier(0.5, 1.6, 0.4, 0.7);
}
.thq-button-animated:hover::before {
  color: var(--dl-color-theme-neutral-light);
  z-index: -1;
  transform: scaleX(1);
}
.Content {
  font-size: 16px;
  font-family: Inter;
  font-weight: 400;
  line-height: 1.15;
  text-transform: none;
  text-decoration: none;
}
.Title {
  font-size: 64px;
  font-style: normal;
  font-family: Inter;
  font-weight: 700px;
  font-stretch: normal;
  letter-spacing: -0.02em;
}
.Subheading {
  font-size: 24px;
  font-style: normal;
  font-family: Inter;
  font-weight: 400px;
  font-stretch: normal;
  letter-spacing: 0em;
}
.Smalltext {
  font-size: 16px;
  font-style: normal;
  font-family: Inter;
  font-weight: 500px;
  font-stretch: normal;
  letter-spacing: 0em;
}
.Heading {
  font-size: 40px;
  font-style: normal;
  font-family: Inter;
  font-weight: 700px;
  font-stretch: normal;
  letter-spacing: 0em;
}
.Bodytext {
  font-size: 20px;
  font-style: normal;
  font-family: Inter;
  font-weight: 500px;
  font-stretch: normal;
  letter-spacing: 0em;
}
@media(max-width: 991px) {
  .thq-grid-4 {
    grid-template-columns: 1fr 1fr 1fr;
  }
}
@media(max-width: 767px) {
  .thq-section-padding {
    padding: var(--dl-space-space-threeunits);
  }
  .thq-flex-column {
    gap: var(--dl-space-space-oneandhalfunits);
  }
  .thq-flex-row {
    gap: var(--dl-space-space-oneandhalfunits);
  }
  .thq-grid-6 {
    grid-gap: var(--dl-space-space-oneandhalfunits);
    grid-template-columns: 1fr 1fr 1fr;
  }
  .thq-grid-5 {
    grid-gap: var(--dl-space-space-oneandhalfunits);
    grid-template-columns: 1fr 1fr 1fr;
  }
  .thq-card {
    padding: var(--dl-space-space-oneandhalfunits);
  }
  .thq-grid-3 {
    grid-gap: var(--dl-space-space-oneandhalfunits);
    grid-template-columns: 1fr 1fr;
  }
  .thq-grid-4 {
    grid-gap: var(--dl-space-space-oneandhalfunits);
    flex-direction: row;
    grid-template-columns: 1fr 1fr;
  }
  .thq-grid-2 {
    grid-gap: var(--dl-space-space-oneandhalfunits);
    grid-template-columns: 1fr;
  }
  .thq-img-scale {
    width: 100%;
  }
}
@media(max-width: 479px) {
  .thq-section-padding {
    padding: var(--dl-space-space-oneandhalfunits);
  }
  .thq-flex-column {
    gap: var(--dl-space-space-unit);
  }
  .thq-flex-row {
    gap: var(--dl-space-space-unit);
  }
  .thq-grid-6 {
    grid-gap: var(--dl-space-space-unit);
    grid-template-columns: 1fr 1fr;
  }
  .thq-grid-5 {
    grid-gap: var(--dl-space-space-unit);
    grid-template-columns: 1fr 1fr;
  }
  .thq-grid-3 {
    grid-gap: var(--dl-space-space-unit);
    align-items: center;
    grid-template-columns: 1fr;
  }
  .thq-grid-4 {
    grid-gap: var(--dl-space-space-unit);
    align-items: center;
    flex-direction: column;
    grid-template-columns: 1fr;
  }
  .thq-grid-2 {
    grid-gap: var(--dl-space-space-unit);
  }
  .thq-grid-auto-300 {
    grid-template-columns: 1fr;
  }
}
</style>

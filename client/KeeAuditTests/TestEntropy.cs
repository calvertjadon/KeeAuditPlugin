using KeeAuditLib.Specifications;
using KeePassLib.Cryptography;
using Microsoft.VisualStudio.TestTools.UnitTesting;

namespace KeeAuditTests
{
    [TestClass]
    public class TestEntropy
    {
        [TestMethod]
        public void TestPassing()
        {
            // arrange
            string password = "very secure password";
            var entropy = QualityEstimation.EstimatePasswordBits(password.ToCharArray());
            var pwEntry = new KeePassLib.PwEntry(true, true);
            pwEntry.Strings.Set("Password", new KeePassLib.Security.ProtectedString(true, password));

            var spec = new EntropySpecification(entropy - 1.0);

            //act
            var result = spec.IsSatisfiedBy(pwEntry);

            //assert
            Assert.AreEqual(true, result);
        }

        [TestMethod]
        public void TestFailing()
        {
            // arrange
            string password = "weak";
            var entropy = QualityEstimation.EstimatePasswordBits(password.ToCharArray());
            var pwEntry = new KeePassLib.PwEntry(true, true);
            pwEntry.Strings.Set("Password", new KeePassLib.Security.ProtectedString(true, password));

            var spec = new EntropySpecification(entropy + 1.0);
            //act
            var result = spec.IsSatisfiedBy(pwEntry);

            //assert
            Assert.AreEqual(false, result);
        }
    }
}

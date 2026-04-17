using KeeAuditLib;
using KeeAuditLib.Specifications;
using KeePassLib.Cryptography;
using Microsoft.VisualStudio.TestTools.UnitTesting;

namespace KeeAuditTests
{
    [TestClass]
    public class TestDuplicatePassword
    {
        [TestMethod]
        public void TestPassing()
        {
            // arrange
            var database = new KeePassLib.PwDatabase
            {
                RootGroup = new KeePassLib.PwGroup(true, true)
            };

            string password = "very secure password";

            var pwEntry1 = new KeePassLib.PwEntry(true, true);
            pwEntry1.Strings.Set("Password", new KeePassLib.Security.ProtectedString(true, password));
            database.RootGroup.AddEntry(pwEntry1, true);

            var repository = new EntryRepository(database);
            var spec = new NotSpecification<KeePassLib.PwEntry>(new DuplicatePasswordSpecification(repository));

            //act
            var result = spec.IsSatisfiedBy(pwEntry1);

            //assert
            Assert.AreEqual(true, result);
        }

        [TestMethod]
        public void TestFailing()
        {
            // arrange
            var database = new KeePassLib.PwDatabase
            {
                RootGroup = new KeePassLib.PwGroup(true, true)
            };

            string password = "very secure password";

            var pwEntry1 = new KeePassLib.PwEntry(true, true);
            pwEntry1.Strings.Set("Password", new KeePassLib.Security.ProtectedString(true, password));
            database.RootGroup.AddEntry(pwEntry1, true);

            var pwEntry2 = new KeePassLib.PwEntry(true, true);
            pwEntry2.Strings.Set("Password", new KeePassLib.Security.ProtectedString(true, password));
            database.RootGroup.AddEntry(pwEntry2, true);

            var repository = new EntryRepository(database);
            var spec = new NotSpecification<KeePassLib.PwEntry>(new DuplicatePasswordSpecification(repository));

            //act
            var result = spec.IsSatisfiedBy(pwEntry1);

            //assert
            Assert.AreEqual(false, result);
        }
    }
}
